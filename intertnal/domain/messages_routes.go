package domain

import (
	"errors"
	"fmt"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"wlcontrol/intertnal/entity"
)

func (a *App) handleMessages(m *tg.Message) (err error) {
	u, err := a.repo.ChatUserState(m.Chat.ID, m.From.ID)

	switch {
	case err == nil:
		err = a.handleStatefulMessage(m, u)
	case errors.Is(err, ErrNotFound):
		err = a.handleStatelessMessage(m)
	}

	return
}

func (a *App) handleStatelessMessage(m *tg.Message) (err error) {
	return
}

func (a *App) handleStatefulMessage(m *tg.Message, u entity.User) (err error) {
	switch u.State {
	case entity.UserStateEnterChatID:
		err = a.handleEnterChatID(m, u)
	}

	return
}

func (a *App) handleEnterChatID(m *tg.Message, u entity.User) (err error) {
	id, err := strconv.ParseInt(m.Text, 10, 64)
	if err != nil {
		return fmt.Errorf("%w: chat id must be a number, not ```%s```", ErrBadRequest, m.Text)
	}

	chat, err := a.repo.ChatByID(id)
	if err != nil {
		if !errors.Is(err, ErrNotFound) {
			return
		}

		chat, err = a.addDefaultChat(id)
	}

	msg := tg.NewMessage(m.Chat.ID, chat.DeviceInfo())

	kb := tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Добавить WL", CallbackAddChatWL),
			tg.NewInlineKeyboardButtonData("Удалить WL", CallbackRemoveChatWL),
		),
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Устройства", CallbackChatDevices),
		),

		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(btnTextBack, CallbackStart),
		),
	)

	msg.ReplyMarkup = &kb

	newMsg, err := a.bot.Send(msg)
	if err != nil {
		return
	}

	a.repo.AddChatUserState(m.Chat.ID, entity.User{
		ID:            m.From.ID,
		ChatID:        m.Chat.ID,
		UserMessageID: m.MessageID,
		BotMessageID:  newMsg.MessageID,
		EditedChatID:  id,
		State:         entity.UserStateEditChatSettings,
	})

	return
}

package domain

import (
	"fmt"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
	"wlcontrol/intertnal/entity"
)

const (
	btnTextBack = "« Назад"
)

const (
	CallbackStart         = "start"
	CallbackChats         = "chats"
	CallbackChat          = "chat"
	CallbackDevices       = "devices"
	CallbackAddChatDevice = "addChatDevice"
	CallbackChatDevices   = "chatDevices"
	CallbackAddChatWL     = "addChatWL"
	CallbackRemoveChatWL  = "removeChatWL"
)

func (a *App) handleCallbacks(cb *tg.CallbackQuery) (err error) {
	_, _ = a.bot.Request(tg.NewCallback(cb.ID, ""))

	s := strings.Split(cb.Data, "_")
	switch len(s) {
	case 1:
		err = a.handleSingleCallbackData(cb)
	case 2:
		err = a.handleMultiCallbackData(cb, s[0], s[1])
	}

	return
}

func (a *App) handleSingleCallbackData(cb *tg.CallbackQuery) (err error) {
	switch cb.Data {
	case CallbackChats:
		err = a.handleCallbackChats(cb)
	case CallbackChat:
		err = a.handleChat(cb)
	case CallbackDevices:
	case CallbackChatDevices:
		err = a.handleCallbackChatDevices(cb)
	case CallbackStart:
		err = a.cmdStart(cb.Message)
	}

	return
}

func (a *App) handleMultiCallbackData(cb *tg.CallbackQuery, data, id string) (err error) {
	_, _ = strconv.ParseInt(id, 10, 64)

	switch data {
	}

	return
}

func (a *App) handleCallbackChats(cb *tg.CallbackQuery) (err error) {
	msg := tg.NewMessage(cb.Message.Chat.ID, "Введите ID чата. Его можно узнать отправив команду `/chatid` в нужном чате.")

	newMsg, err := a.bot.Send(msg)
	if err != nil {
		return
	}

	a.repo.AddChatUserState(cb.Message.Chat.ID, entity.User{
		ID:            cb.From.ID,
		ChatID:        cb.Message.Chat.ID,
		UserMessageID: cb.Message.MessageID,
		BotMessageID:  newMsg.MessageID,
		State:         entity.UserStateEnterChatID,
	})

	return
}

func (a *App) handleCallbackChatDevices(cb *tg.CallbackQuery) error {
	m := cb.Message

	u, err := a.repo.ChatUserState(m.Chat.ID, cb.From.ID)
	if err != nil {
		return fmt.Errorf("user state: %s", err)
	}

	_, err = a.repo.ChatByID(u.EditedChatID)
	if err != nil {
		return fmt.Errorf("chat by id: %s", err)

	}

	devices, err := a.repo.ChatDevices(u.EditedChatID)
	if err != nil {
		return fmt.Errorf("chat devices: %s", err)

	}

	btnRows := make([][]tg.InlineKeyboardButton, 0, len(devices)+1)

	for _, v := range devices {
		btnData := fmt.Sprintf("%s_%d", CallbackAddChatDevice, v.ID)
		btnRows = append(btnRows, tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData(v.Name, btnData)))
	}

	btnRows = append(btnRows, tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData(btnTextBack, CallbackChat)))

	msg := tg.NewEditMessageText(m.Chat.ID, u.BotMessageID, "Выберите чат, который нужно добавить или удалить в данный чат")

	kb := tg.NewInlineKeyboardMarkup(btnRows...)
	msg.ReplyMarkup = &kb

	newMsg, err := a.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("send response: %s", err)
	}

	u.BotMessageID = newMsg.MessageID
	a.repo.AddChatUserState(m.Chat.ID, u)

	return nil
}

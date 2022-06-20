package domain

import (
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
	case CallbackDevices:
	case CallbackChatDevices:
		err = a.handleCallbackChatDevices(cb.Message)
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
	msg := tg.NewMessage(cb.Message.Chat.ID, "Укажите ID чата. Его можно узнать отправив команду `/chatid` в нужном вам чате.")

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

func (a *App) handleCallbackChatDevices(m *tg.Message) (err error) {

	return
}

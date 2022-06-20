package domain

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"wlcontrol/intertnal/entity"
)

func defaultDevices(chatID int64) []entity.Mikrotik {
	return []entity.Mikrotik{
		{
			ID:     1,
			ChatID: chatID,
			WL:     "WL",
		},
		{
			ID:     2,
			ChatID: chatID,
			WL:     "WL",
		},
	}
}

var keyboardChats = tg.NewInlineKeyboardMarkup(
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

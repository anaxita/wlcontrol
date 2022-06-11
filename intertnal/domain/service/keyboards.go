package service

import tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const (
	btnStart = "start"

	btnRouters     = "routers"
	btnAddRouter   = "add router"
	btnShowRouters = "show routers"

	btnChats          = "chats"
	btnChat           = "chat"
	btnEditChatWL     = "set wl"
	btnSetChatDevices = "setChatDevices"
	btnSetChatDevice  = "setChatDevice"
	btnChangeDeviceWL = "changeChatDevice"
)

const (
	textBtnBack = "« Назад"

	textBtnRouters     = "Микротики"
	textBtnAddRouter   = "Добавить микротик"
	textBtnShowRouters = "Показать все"

	textBtnChats          = "Чаты"
	textBtnEditChatWL     = "Изменить WL"
	textBtnSetChatDevices = "Задать микротики"
)

var kbStart = tg.NewInlineKeyboardMarkup(
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(textBtnRouters, btnRouters),
	),
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(textBtnChats, btnChats),
	),
)

var kbChats = tg.NewInlineKeyboardMarkup(
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(textBtnEditChatWL, btnEditChatWL),
		tg.NewInlineKeyboardButtonData(textBtnSetChatDevices, btnSetChatDevices),
	),

	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(textBtnBack, btnStart),
	),
)

var kbRouters = tg.NewInlineKeyboardMarkup(
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(textBtnAddRouter, btnAddRouter),
		tg.NewInlineKeyboardButtonData(textBtnShowRouters, btnShowRouters),
	),

	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(textBtnBack, btnStart),
	),
)

var kbAddRouter = tg.NewInlineKeyboardMarkup(
	tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(textBtnBack, btnRouters),
	),
)

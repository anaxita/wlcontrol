package service

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"wlcontrol/intertnal/domain/entity"
)

func (c *Core) cbStart(cb *tg.CallbackQuery) error {
	m := cb.Message

	msg := tg.NewEditMessageText(m.Chat.ID, m.MessageID, "Выберите раздел:")
	msg.ReplyMarkup = &kbStart

	_, err := c.bot.Send(msg)

	return err
}

func (c *Core) cbChats(cb *tg.CallbackQuery) error {
	m := cb.Message

	msg := tg.NewEditMessageText(m.Chat.ID, m.MessageID, "Выберите операцию с чатами:")
	msg.ReplyMarkup = &kbChats

	_, err := c.bot.Send(msg)
	if err != nil {
		return err
	}

	return err
}

func (c *Core) cbRouters(cb *tg.CallbackQuery) error {
	m := cb.Message

	msg := tg.NewEditMessageText(m.Chat.ID, m.MessageID, "Выберите операцию с микротиками:")
	msg.ReplyMarkup = &kbRouters

	_, err := c.bot.Send(msg)
	if err != nil {
		return err
	}

	return err
}

func (c *Core) cbAddRouter(cb *tg.CallbackQuery) error {
	m := cb.Message

	msg := tg.NewMessage(m.Chat.ID, textAddRouter)

	_, err := c.bot.Send(msg)
	if err != nil {
		return err
	}

	c.repo.AddChatUser(m.Chat.ID, entity.User{
		ID:        cb.From.ID,
		MessageID: m.MessageID,
		State:     entity.UserStateAddRouter,
	})

	return err
}

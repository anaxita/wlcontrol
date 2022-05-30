package service

import (
	"errors"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"wlcontrol/intertnal/domain"
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
	chatID := cb.Message.Chat.ID

	msg := tg.NewMessage(chatID, "Введите id чата")
	_, err := c.bot.Send(msg)
	if err != nil {
		return err
	}

	c.repo.AddChatUserState(chatID, entity.User{
		ID:        cb.From.ID,
		MessageID: cb.Message.MessageID,
		ChatID:    chatID,
		State:     entity.UserStateEnterChatID,
	})

	return nil
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

	c.repo.AddChatUserState(m.Chat.ID, entity.User{
		ID:        cb.From.ID,
		MessageID: cb.Message.MessageID,
		ChatID:    m.Chat.ID,
		State:     entity.UserStateAddRouter,
	})

	return err
}

func (c *Core) cbEditChatWL(cb *tg.CallbackQuery) error {
	m := cb.Message

	// TODO: use chat
	chat, err := c.repo.ChatByID(m.Chat.ID)
	switch {
	case err == nil:
	case !errors.Is(err, domain.ErrNotFound):
		return err
	default:
	}

	msg := tg.NewMessage(m.Chat.ID, textAddRouter)

	_, err = c.bot.Send(msg)
	if err != nil {
		return err
	}

	return err
}

package service

import (
	"errors"
	"fmt"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"wlcontrol/intertnal/domain"
	"wlcontrol/intertnal/domain/entity"
)

func (c *Core) msgAddRouter(m *tg.Message) error {
	r, err := parseRouter(m.Text)
	if err != nil {
		return fmt.Errorf("parse router data: %w", err)
	}

	err = c.repo.AddRouter(r)
	if err != nil {
		return fmt.Errorf("add router: %w", err)
	}

	msg := tg.NewMessage(m.Chat.ID, "Микротик успешно добавлен!")
	msg.ReplyMarkup = &kbAddRouter

	_, err = c.bot.Send(msg)

	return err
}

func (c *Core) msgSetDeviceToChat(m *tg.Message) error {
	msg := tg.NewMessage(m.Chat.ID, "Пожалуйста, выберите микротик, который нужно задать данному чату")
	msg.ReplyMarkup = &kbRouters

	_, err := c.bot.Send(msg)

	return err
}

func (c *Core) msgShowChatSettings(m *tg.Message) error {
	id, err := strconv.ParseInt(m.Text, 10, 64)
	if err != nil {
		return errors.New("айди чата должен быть числом")
	}

	var text string

	chat, err := c.repo.ChatByID(id)
	switch {
	case err == nil:
		text = chat.Info()
	case !errors.Is(err, domain.ErrNotFound):
		return fmt.Errorf("не смогли найти такой чат: %w", err)
	default:
		text = "У данного чата стандартные настройки. Выберите, что хотите изменить:"
	}

	msg := tg.NewMessage(m.Chat.ID, text)
	msg.ReplyMarkup = &kbChats

	_, err = c.bot.Send(msg)
	if err != nil {
		return err
	}

	c.repo.AddChatUserState(m.Chat.ID, entity.User{
		ID:        m.From.ID,
		MessageID: m.MessageID,
		ChatID:    id,
		State:     entity.UserStateEditChatSettings,
	})

	return err
}

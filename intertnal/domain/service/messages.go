package service

import (
	"errors"
	"fmt"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
	"wlcontrol/intertnal/domain"
	"wlcontrol/intertnal/domain/entity"
)

func (c *App) msgAddRouter(m *tg.Message) error {
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

func (c *App) msgSetDeviceToChat(m *tg.Message, u entity.User) error {
	s := strings.Split(m.Text, "_")
	if len(s) != 2 {
		return fmt.Errorf("%w: undefined device id in string '%s'", domain.ErrBadRequest, s)
	}

	id, err := strconv.ParseInt(s[1], 10, 64)
	if err != nil {
		return fmt.Errorf("%w: device id must be a string, got: '%s'", domain.ErrBadRequest, s[1])
	}

	device, err := c.repo.DeviceByID(id)
	if err != nil {
		return err
	}

	device.ChatID = u.EditedChatID

	chat, err := c.repo.ChatByID(u.EditedChatID)
	if err != nil {
		return err
	}

	if _, ok := chat.IsDeviceFound(id); ok {
		err = c.repo.RemoveDeviceFromChat(device)
	} else {
		err = c.repo.AddDevicesToChat(device)
	}

	if err != nil {
		return err
	}

	return c.cbEditChatDevices(m, u)
}

func (c *App) msgShowChatSettings(m *tg.Message) error {
	id, err := strconv.ParseInt(m.Text, 10, 64)
	if err != nil {
		return errors.New("айди чата должен быть числом")
	}

	var text string

	chat, err := c.repo.ChatByID(id)
	switch {
	case err == nil:
	case !errors.Is(err, domain.ErrNotFound):
		return fmt.Errorf("не смогли найти такой чат: %w", err)
	default:
		devices := []entity.Mikrotik{
			{
				ID:     1,
				ChatID: id,
				WL:     "wl",
			},
			{
				ID:     2,
				ChatID: id,
				WL:     "wl",
			},
		}
		err = c.repo.AddDevicesToChat(devices...)
		if err != nil {
			return err
		}

		chat, err = c.repo.ChatByID(id)
		if err != nil {
			return err
		}
	}

	text = chat.Info()

	msg := tg.NewMessage(m.Chat.ID, text)
	msg.ReplyMarkup = &kbChats

	_, err = c.bot.Send(msg)
	if err != nil {
		return err
	}

	c.repo.AddChatUserState(m.Chat.ID, entity.User{
		ID:           m.From.ID,
		MessageID:    m.MessageID,
		ChatID:       m.Chat.ID,
		EditedChatID: id,
		State:        entity.UserStateEditChatSettings,
	})

	return err
}

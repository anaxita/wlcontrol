package service

import (
	"errors"
	"fmt"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"wlcontrol/intertnal/domain"
	"wlcontrol/intertnal/domain/entity"
)

func (c *App) cbStart(cb *tg.CallbackQuery) error {
	m := cb.Message

	msg := tg.NewEditMessageText(m.Chat.ID, m.MessageID, "Выберите раздел:")
	msg.ReplyMarkup = &kbStart

	_, err := c.bot.Send(msg)

	return err
}

func (c *App) cbChats(cb *tg.CallbackQuery) error {
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

func (c *App) cbChat(cb *tg.CallbackQuery, u entity.User) error {
	m := cb.Message

	chat, err := c.repo.ChatByID(u.EditedChatID)
	if err != nil {
		return err
	}

	msg := tg.NewEditMessageText(m.Chat.ID, m.MessageID, chat.Info())
	msg.ReplyMarkup = &kbChats

	_, err = c.bot.Send(msg)

	return err
}

func (c *App) cbRouters(cb *tg.CallbackQuery) error {
	m := cb.Message

	msg := tg.NewEditMessageText(m.Chat.ID, m.MessageID, "Выберите операцию с микротиками:")
	msg.ReplyMarkup = &kbRouters

	_, err := c.bot.Send(msg)
	if err != nil {
		return err
	}

	return err
}

func (c *App) cbAddRouter(cb *tg.CallbackQuery) error {
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

func (c *App) cbEditChatWL(cb *tg.CallbackQuery, u entity.User) error {
	m := cb.Message

	chat, err := c.repo.ChatByID(u.EditedChatID)
	switch {
	case err == nil:
	case !errors.Is(err, domain.ErrNotFound):
		return err
	default:
	}

	var rows [][]tg.InlineKeyboardButton
	for _, v := range chat.Devices {
		text := fmt.Sprintf("%s (%s)", v.Name, v.WL)
		btn := tg.NewInlineKeyboardButtonData(text, fmt.Sprintf("%s_%d", btnChangeDeviceWL, v.ID))
		row := tg.NewInlineKeyboardRow(btn)
		rows = append(rows, row)
	}

	rows = append(rows, tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData(textBtnBack, btnChat)))
	kb := tg.NewInlineKeyboardMarkup(rows...)

	msg := tg.NewEditMessageText(m.Chat.ID, cb.Message.MessageID, "Выберите микротик, на котором нужно изменить WL:")
	msg.ReplyMarkup = &kb

	_, err = c.bot.Send(msg)
	if err != nil {
		return err
	}

	return err
}

func (c *App) cbEditChatDevices(m *tg.Message, u entity.User) error {
	chat, err := c.repo.ChatByID(u.EditedChatID)
	switch {
	case err == nil:
	case !errors.Is(err, domain.ErrNotFound):
		return fmt.Errorf("не смогли найти такой чат: %w", err)
	default:
		devices := []entity.Mikrotik{
			{
				ID:     1,
				ChatID: u.EditedChatID,
				WL:     "wl",
			},
			{
				ID:     2,
				ChatID: u.EditedChatID,
				WL:     "wl",
			},
		}
		err = c.repo.AddDevicesToChat(devices...)
		if err != nil {
			return err
		}

		chat.Devices = devices
	}

	devices, err := c.repo.Devices()
	if err != nil {
		return err
	}

	var rows [][]tg.InlineKeyboardButton
	for _, v := range devices {
		text := v.Name

		if _, ok := chat.IsDeviceFound(v.ID); ok {
			text = "✔ " + text
		}

		btn := tg.NewInlineKeyboardButtonData(text, fmt.Sprintf("%s_%d", btnSetChatDevice, v.ID))
		row := tg.NewInlineKeyboardRow(btn)
		rows = append(rows, row)
	}

	rows = append(rows, tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData(textBtnBack, btnChat)))
	kb := tg.NewInlineKeyboardMarkup(rows...)

	text := "Пожалуйста, выберите микротики, которые нужно добавить/удалить у данного чата"

	msg := tg.NewEditMessageText(m.Chat.ID, m.MessageID, text)
	msg.ReplyMarkup = &kb

	_, err = c.bot.Send(msg)
	if err != nil {
		return err
	}

	// update cache
	u.MessageID = m.MessageID
	u.State = entity.UserStateSetDeviceToChat

	c.repo.AddChatUserState(m.Chat.ID, u)

	return nil
}

func (c *App) cbEditChatDeviceWL(m *tg.Message, u entity.User) error {
	device, err := c.repo.DeviceByID(u.MikrotikID)
	if err != nil {
		return err
	}

	device.ChatID = u.EditedChatID
	device.WL = "wl" // default address list name

	chat, err := c.repo.ChatByID(u.EditedChatID)
	if err != nil {
		return err
	}

	if _, ok := chat.IsDeviceFound(u.MikrotikID); ok {
		err = c.repo.RemoveDeviceFromChat(device)
	} else {
		err = c.repo.AddDevicesToChat(device)
	}

	if err != nil {
		return err
	}

	return c.cbEditChatDevices(m, u)
}

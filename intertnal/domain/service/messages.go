package service

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Core) msgAddRouter(m *tg.Message) error {
	u, err := c.repo.ChatUser(m.Chat.ID, m.From.ID)
	if err != nil {
		return err
	}

	_, err = parseRouter(m.Text)
	if err != nil {
		return err
	}

	c.repo.DeleteChatUser(m.Chat.ID, m.From.ID)

	msg := tg.NewMessage(m.Chat.ID, "Микротик успешно добавлен!")
	msg.ReplyToMessageID = u.MessageID
	msg.ReplyMarkup = &kbAddRouter

	_, err = c.bot.Send(msg)

	return err
}

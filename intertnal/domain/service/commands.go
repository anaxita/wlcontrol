package service

import tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const (
	cmdStart = "start"
)

func (c *Core) cmdStart(m *tg.Message) error {
	msg := tg.NewMessage(m.Chat.ID, "Выберите раздел:")
	msg.ReplyMarkup = &kbStart

	_, err := c.bot.Send(msg)

	return err
}

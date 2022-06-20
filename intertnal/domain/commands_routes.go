package domain

import tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const (
	CommandStart = "start"
)

func (a *App) handleCommands(m *tg.Message) error {
	var err error

	switch m.Command() {
	case CommandStart:
		err = a.cmdStart(m)
	}

	return err
}

func (a *App) cmdStart(m *tg.Message) error {
	msg := tg.NewMessage(m.Chat.ID, "Выберите действите")

	kb := tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Микротики", CallbackDevices),
		),
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("Чаты", CallbackChats),
		),
	)

	msg.ReplyMarkup = &kb

	_, err := a.bot.Send(msg)

	return err
}

package service

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *Core) handleCommand(m *tg.Message) {
	if m.From.IsBot {
		return
	}

	var err error

	switch m.Command() {
	case cmdStart:
		err = c.cmdStart(m)
	}

	if err != nil {
		log.Println("handleCommand: ", err)
	}
}

func (c *Core) handleMessage(m *tg.Message) {
	if m.From.IsBot {
		return
	}

	var err error

	err = c.msgAddRouter(m)

	if err != nil {
		_, _ = c.bot.Send(tg.NewMessage(m.Chat.ID, err.Error()))
		log.Println("handleMessage: ", err)
	}
}

func (c *Core) handleCallback(cb *tg.CallbackQuery) {
	_, err := c.bot.Request(tg.NewCallback(cb.ID, ""))
	if err != nil {
		log.Println("[CALLBACK REQUEST] ", err)
		return
	}

	if cb.From.IsBot {
		return
	}

	switch cb.Data {
	case btnChats:
		err = c.cbChats(cb)
	case btnStart:
		err = c.cbStart(cb)
	case btnRouters:
		err = c.cbRouters(cb)
	case btnAddRouter:
		err = c.cbAddRouter(cb)
	}

	if err != nil {
		log.Println("handleCallback: ", err)
	}
}

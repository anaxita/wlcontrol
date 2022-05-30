package service

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"wlcontrol/intertnal/domain/entity"
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

	u, err := c.repo.ChatUserState(m.Chat.ID, m.From.ID)
	if err != nil {
		err = c.handeDefaultMessage(m)
	} else {
		err = c.handeStateMessage(m, u)
	}
	if err != nil {
		_, _ = c.bot.Send(tg.NewMessage(m.Chat.ID, err.Error()))
		log.Println("handleMessage: ", err)
	}
}

func (c *Core) handeStateMessage(m *tg.Message, u entity.User) (err error) {
	switch u.State {
	case entity.UserStateAddRouter:
		err = c.msgAddRouter(m)
	case entity.UserStateSetDeviceToChat:
		err = c.msgSetDeviceToChat(m)
	case entity.UserStateEnterChatID:
		err = c.msgShowChatSettings(m)
	}

	if err != nil {
		return err
	}

	c.repo.DeleteChatUserState(m.Chat.ID, m.From.ID)
	return nil
}
func (c *Core) handeDefaultMessage(m *tg.Message) error {
	return nil
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
	case btnEditChatWL:
		err = c.cbEditChatWL(cb)
	}

	if err != nil {
		log.Println("handleCallback: ", err)
	}
}

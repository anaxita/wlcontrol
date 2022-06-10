package service

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"wlcontrol/intertnal/domain/entity"
)

func (c *App) handleCommand(m *tg.Message) {
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

func (c *App) handleMessage(m *tg.Message) {
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

func (c *App) handeStateMessage(m *tg.Message, u entity.User) (err error) {
	switch u.State {
	case entity.UserStateAddRouter:
		err = c.msgAddRouter(m)
	case entity.UserStateEnterChatID:
		err = c.msgShowChatSettings(m)
	}

	return err
}
func (c *App) handeDefaultMessage(m *tg.Message) error {
	return nil
}

func (c *App) handleCallback(cb *tg.CallbackQuery) {
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
	case btnChat:
		err = c.cbChat(cb)
	case btnStart:
		err = c.cbStart(cb)
	case btnRouters:
		err = c.cbRouters(cb)
	case btnAddRouter:
		err = c.cbAddRouter(cb)
	case btnEditChatWL:
		u, err := c.repo.ChatUserState(cb.Message.Chat.ID, cb.From.ID)
		if err != nil {
			break
		}

		err = c.cbEditChatWL(cb, u.ChatID)
	case btnSetChatDevices:
		u, err := c.repo.ChatUserState(cb.Message.Chat.ID, cb.From.ID)
		if err != nil {
			break
		}

		err = c.cbEditChatDevices(cb.Message, u)
	default:
		u, err := c.repo.ChatUserState(cb.Message.Chat.ID, cb.From.ID)
		if err != nil {
			break
		}

		cb.Message.Text = cb.Data
		err = c.msgSetDeviceToChat(cb.Message, u)
	}

	if err != nil {
		log.Println("handleCallback: ", err)
	}
}

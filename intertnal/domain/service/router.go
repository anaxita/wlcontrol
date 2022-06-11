package service

import (
	"fmt"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
	"wlcontrol/intertnal/domain"
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
		u, err := c.repo.ChatUserState(cb.Message.Chat.ID, cb.From.ID)
		if err != nil {
			break
		}

		err = c.cbChat(cb, u)
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

		err = c.cbEditChatWL(cb, u)
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

		err = c.multiCallback(cb, u)
	}

	if err != nil {
		log.Println("handleCallback: ", err)
	}
}

func (c *App) multiCallback(cb *tg.CallbackQuery, u entity.User) (err error) {
	m := cb.Message

	s := strings.Split(m.Text, "_")
	if len(s) != 2 {
		return fmt.Errorf("%w: undefined device id in string '%s'", domain.ErrBadRequest, s)
	}

	id, err := strconv.ParseInt(s[1], 10, 64)
	if err != nil {
		return fmt.Errorf("%w: device id must be a string, got: '%s'", domain.ErrBadRequest, s[1])
	}

	u.MikrotikID = id

	switch s[0] {
	case btnChangeDeviceWL:
		err = c.cbEditChatWL(cb, u)
	case btnSetChatDevice:
		err = c.msgSetDeviceToChat(cb.Message, u)
	}

	return err
}

package service

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"wlcontrol/intertnal/infrastructure/dal"
	"wlcontrol/intertnal/infrastructure/mikrotik"
)

type App struct {
	bot    *tg.BotAPI
	repo   *dal.Repository
	device *mikrotik.Device
}

func New(botToken string, isDebug bool, repo *dal.Repository, device *mikrotik.Device) (App, error) {
	bot, err := tg.NewBotAPI(botToken)
	if err != nil {
		return App{}, err
	}
	bot.Debug = isDebug

	return App{
		bot:    bot,
		repo:   repo,
		device: device,
	}, nil
}

func (c *App) Run() {
	defer c.bot.StopReceivingUpdates()

	u := tg.NewUpdate(0)
	u.Timeout = 60

	updates := c.bot.GetUpdatesChan(u)
	for u := range updates {
		go c.handle(u)
	}
}

func (c *App) handle(u tg.Update) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("[PANIC] ", err)
		}
	}()

	switch {
	case u.CallbackQuery != nil:
		c.handleCallback(u.CallbackQuery)
	case u.Message == nil || u.Message.From.IsBot: // ignoring empty messages and messages from bots
	// case u.Message.Chat.IsPrivate(): TODO: uncomment before release
	// 	_, _ = c.bot.Send(tg.NewMessage(u.Message.Chat.ID, textPrivateStart))  TODO: uncomment before release
	case u.Message.IsCommand():
		c.handleCommand(u.Message)
	default:
		c.handleMessage(u.Message)
	}
}

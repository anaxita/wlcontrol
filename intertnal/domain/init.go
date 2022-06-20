package domain

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"wlcontrol/intertnal/entity"
	"wlcontrol/intertnal/infrastructure/mikrotik"
)

type Repository interface {
	ChatUserState(chatID, userID int64) (entity.User, error)
	AddChatUserState(chatID int64, u entity.User)
	DeleteChatUserState(chatID, userID int64)
	AddRouter(router entity.MikrotikCreate) error
	ChatByID(id int64) (entity.Chat, error)
	ChatDevices(id int64) ([]entity.Mikrotik, error)
	AddDevicesToChat(devices ...entity.Mikrotik) error
	RemoveDeviceFromChat(devices ...entity.Mikrotik) error
	DeviceByID(id int64) (device entity.Mikrotik, err error)
	Devices() (devices []entity.Mikrotik, err error)
}

type App struct {
	bot    *tg.BotAPI
	repo   Repository
	device *mikrotik.Device
}

func New(botToken string, isDebug bool, repo Repository, device *mikrotik.Device) (App, error) {
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

func (a *App) Run() {
	defer a.bot.StopReceivingUpdates()

	u := tg.NewUpdate(0)
	u.Timeout = 60

	updates := a.bot.GetUpdatesChan(u)
	for u := range updates {
		a.handle(u)
	}
}

func (a *App) handle(u tg.Update) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("[PANIC] ", err)
		}
	}()

	var err error

	switch {
	case u.CallbackQuery != nil:
		err = a.handleCallbacks(u.CallbackQuery)
	case u.Message == nil || u.Message.From.IsBot: // ignore empty messages and messages from bots
	case u.Message.IsCommand():
		err = a.handleCommands(u.Message)
	default:
		err = a.handleMessages(u.Message)
	}

	if err != nil {
		log.Println("[HANDLE] ", err)
	}
}

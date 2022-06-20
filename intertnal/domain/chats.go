package domain

import "wlcontrol/intertnal/entity"

func (a *App) addDefaultChat(chatID int64) (chat entity.Chat, err error) {
	err = a.repo.AddDevicesToChat(defaultDevices(chatID)...)
	if err != nil {
		return
	}

	chat, err = a.repo.ChatByID(chatID)

	return
}

package domain

import "wlcontrol/intertnal/entity"

func defaultDevices(chatID int64) []entity.Mikrotik {
	return []entity.Mikrotik{
		{
			ID:     1,
			ChatID: chatID,
			WL:     "WL",
		},
		{
			ID:     2,
			ChatID: chatID,
			WL:     "WL",
		},
	}
}

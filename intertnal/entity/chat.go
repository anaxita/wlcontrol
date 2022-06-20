package entity

import "fmt"

type Chat struct {
	ID           int64
	AddressLists []Mikrotik
}

func (c Chat) DeviceInfo() string {
	var text string
	for i, v := range c.AddressLists {
		text += fmt.Sprintf("\n%d. %s --> %s", i+1, v.Name, v.WL)
	}
	return "Этот чат имеет следующие настройки (Mikrotik -> WL):\n" + text
}

func (c Chat) IsDeviceFound(deviceID int64) (Mikrotik, bool) {
	for _, v := range c.AddressLists {
		if v.ID == deviceID {
			return v, true
		}
	}

	return Mikrotik{}, false
}

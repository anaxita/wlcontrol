package entity

import "fmt"

type Chat struct {
	ID      int64
	Devices []Mikrotik
}

func (c Chat) Info() string {
	var text string
	for i, v := range c.Devices {
		text += fmt.Sprintf("\n%d. %s --> %s", i+1, v.Name, v.WL)
	}
	return "Этот чат имеет следующие настройки (Mikrotik -> WL):\n" + text
}

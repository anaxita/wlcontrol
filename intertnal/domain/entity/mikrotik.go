package entity

import (
	"errors"
)

type Mikrotik struct {
	ID       int64  `db:"id"`
	ChatID   int64  `db:"chat_id"`
	Name     string `db:"name"`
	Address  string `db:"address"`
	Login    string `db:"login"`
	Password string `db:"password"`
	WL       string `db:"wl"`
}

type MikrotikCreate struct {
	Name     string `db:"name"`
	Address  string `db:"address"`
	Login    string `db:"login"`
	Password string `db:"password"`
}

func (m MikrotikCreate) Validate() error {
	var errText string

	switch {
	case len(m.Name) == 0:
		errText += "Name is empty, "
		fallthrough
	case len(m.Address) == 0:
		errText += "Address is empty, "
		fallthrough
	case len(m.Login) == 0:
		errText += "Login is empty, "
		fallthrough
	case len(m.Password) == 0:
		errText += "Password is empty, "
	}

	if len(errText) != 0 {
		errText = errText[:len(errText)-2] // trim 2 last signs
		return errors.New(errText)
	}

	return nil
}

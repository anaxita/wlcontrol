package entity

import (
	"errors"
)

type Mikrotik struct {
	ID       int64
	Name     string
	Address  string
	Login    string
	Password string

	ChatID ChatID
	WL     string
}

type MikrotikCreate struct {
	Name     string
	Address  string
	Login    string
	Password string
}

func (m MikrotikCreate) Validate() error {
	var errText string

	switch {
	case len(m.Name) == 0:
		errText += " Name is empty,"
		fallthrough
	case len(m.Address) == 0:
		errText += " Address is empty,"
		fallthrough
	case len(m.Login) == 0:
		errText += " Login is empty,"
		fallthrough
	case len(m.Password) == 0:
		errText += " Password is empty,"
	}

	if len(errText) != 0 {
		errText = errText[:len(errText)-2] // trim last sign
		return errors.New(errText)
	}

	return nil
}

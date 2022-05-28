package service

import (
	"fmt"
	"strings"
	"wlcontrol/intertnal/domain/entity"
)

func parseRouter(text string) (entity.MikrotikCreate, error) {
	s := strings.Split(text, "\n")
	if len(s) != 4 {
		return entity.MikrotikCreate{}, fmt.Errorf("нужно прислать 4 строки, а вы прислали %d", len(s))
	}

	m := entity.MikrotikCreate{
		Name:     s[0],
		Address:  s[1],
		Login:    s[2],
		Password: s[3],
	}

	return m, m.Validate()
}

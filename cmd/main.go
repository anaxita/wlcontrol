package main

import (
	"log"
	"wlcontrol/intertnal/bootstrap"
	"wlcontrol/intertnal/domain/service"
	"wlcontrol/intertnal/infrastructure/dal"
	"wlcontrol/intertnal/infrastructure/mikrotik"
)

func main() {
	c, err := bootstrap.NewConfig()
	if err != nil {
		log.Fatal("config: ", err)
	}

	repo := dal.NewRepository()
	device := mikrotik.NewDevice()

	core, err := service.NewCore(c.BotToken, c.BotDebug, repo, device)
	if err != nil {
		log.Fatal("core: ", err)
	}

	core.Run()
}

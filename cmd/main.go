package main

import (
	_ "github.com/mattn/go-sqlite3"
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

	repo, err := dal.NewRepository(c.DBName)
	if err != nil {
		log.Fatal("repository: ", err)
	}

	device := mikrotik.New()

	core, err := service.New(c.BotToken, c.BotDebug, repo, device)
	if err != nil {
		log.Fatal("core: ", err)
	}

	core.Run()
}

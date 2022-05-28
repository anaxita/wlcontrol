package bootstrap

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	HTTPPort string
	BotToken string
	BotDebug bool
}

func NewConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, fmt.Errorf("load .env: %s", err)
	}

	var c Config

	isDebug, err := strconv.ParseBool(os.Getenv("BOT_DEBUG"))
	if err != nil {
		return Config{}, fmt.Errorf("parse BOT_DEBUG: %s", err)
	}

	c.HTTPPort = os.Getenv("HTTP_PORT")
	c.BotToken = os.Getenv("BOT_TOKEN")
	c.BotDebug = isDebug

	return c, nil
}

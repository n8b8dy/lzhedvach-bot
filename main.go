package main

import (
	"github.com/joho/godotenv"
	"github.com/n8b8dy/lzhedvach-bot/functionality/information"
	"github.com/n8b8dy/lzhedvach-bot/functionality/lzhedvacher"
	tele "gopkg.in/telebot.v3"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		return
	}

	pref := tele.Settings{
		Token: os.Getenv("BOT_TOKEN"),
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Groups
	information.CreateGroup(b)
	lzhedvacher.CreateGroup(b)

	b.Start()
}

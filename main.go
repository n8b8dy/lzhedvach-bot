package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"

	"github.com/n8b8dy/lzhedvach-bot/functionality/information"
	"github.com/n8b8dy/lzhedvach-bot/functionality/lzhedvacher"
	"github.com/n8b8dy/lzhedvach-bot/functionality/utils"
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
	utils.CreateGroup(b)

	b.Start()
}

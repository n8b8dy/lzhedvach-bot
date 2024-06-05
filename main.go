package main

import (
	"fmt"
	"github.com/joho/godotenv"
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

	b.Handle("/start", func(c tele.Context) error {
		text := fmt.Sprintf("Hi, %s! This is a bot, created by @n8body, specially for <b>Lzhedvach</b> group.", c.Message().Sender.FirstName)

		return c.Reply(text, &tele.SendOptions{
			ParseMode: tele.ModeHTML,
		})
	})

	b.Start()
}

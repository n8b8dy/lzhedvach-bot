package lzhedvacher

import (
	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
	"log"
	"os"
	"strconv"
)

func CreateGroup(b *tele.Bot) {
	lzhedvachID, err := strconv.ParseInt(os.Getenv("LZHEDVACH_ID"), 10, 64)
	if err != nil {
		log.Fatal(err)
		return
	}

	group := b.Group()
	group.Use(middleware.Whitelist(lzhedvachID))
}

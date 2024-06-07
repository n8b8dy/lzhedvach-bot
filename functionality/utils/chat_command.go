package utils

import (
	"fmt"

	tele "gopkg.in/telebot.v3"
)

func chatCommandHandler(c tele.Context) error {
	text := fmt.Sprintf("ID: <code>%d</code>", c.Chat().ID)

	return c.Reply(text, &tele.SendOptions{
		ParseMode: tele.ModeHTML,
	})
}

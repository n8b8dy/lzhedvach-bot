package information

import (
	"fmt"
	tele "gopkg.in/telebot.v3"
)

func startCommandHandler(c tele.Context) error {
	text := fmt.Sprintf("Hi, %s! This is a bot, created by @n8body, specially for <b>Lzhedvach</b> group.", c.Message().Sender.FirstName)

	return c.Reply(text, &tele.SendOptions{
		ParseMode: tele.ModeHTML,
	})
}

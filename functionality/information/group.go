package information

import tele "gopkg.in/telebot.v3"

func CreateGroup(b *tele.Bot) {
	group := b.Group()

	group.Handle("/start", startCommandHandler)
}

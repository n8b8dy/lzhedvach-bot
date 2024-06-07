package lzhedvacher

import (
	"os"

	"github.com/sashabaranov/go-openai"
	tele "gopkg.in/telebot.v3"
)

func CreateGroup(b *tele.Bot) {
	group := b.Group()
	group.Use(AllowLzhedvachOnly())

	client := openai.NewClient(os.Getenv("OPENAI_KEY"))

	group.Handle(tele.OnText, createChatterTextHandler(client))
}

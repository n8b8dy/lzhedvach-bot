package lzhedvacher

import (
	"context"
	"math/rand/v2"

	"github.com/sashabaranov/go-openai"
	tele "gopkg.in/telebot.v3"
)

func createChatterTextHandler(client *openai.Client) tele.HandlerFunc {
	return func(c tele.Context) error {
		// 5% chance to reply
		if rand.IntN(100) > 5 {
			return nil
		}

		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT4o,
				Messages: []openai.ChatCompletionMessage{
					{
						Role: openai.ChatMessageRoleSystem,
						Content: "You are a funny strange man in a Telegram group, who tries to convince everyone that he is another man from the group named Лё. " +
							"Respond to users' messages in a quirky and funny manner. Be unpredictable as possible!",
					},
					{
						Role:    openai.ChatMessageRoleUser,
						Name:    c.Sender().FirstName,
						Content: c.Text(),
					},
				},
			},
		)

		if err != nil {
			return nil
		}

		return c.Reply(resp.Choices[0].Message.Content)
	}
}

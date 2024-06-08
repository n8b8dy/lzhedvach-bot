package lzhedvacher

import (
	"fmt"
	"math/rand/v2"
	"regexp"
	"strings"

	tele "gopkg.in/telebot.v3"
)

func daTextHandler(c tele.Context) error {
	regex := regexp.MustCompile("\\s*([dд]+[aа]+[\\s.,!?\\\\()01:;\"']*)$")

	match := regex.FindString(strings.ToLower(c.Text()))
	if len(match) == 0 {
		return nil
	}

	// 10% chance to send
	if rand.IntN(100) > 10 {
		return nil
	}

	return c.Reply(fmt.Sprintf("Пиз%s", match))
}

package lzhedvacher

import (
	"log"
	"os"
	"strconv"

	tele "gopkg.in/telebot.v3"
)

func AllowLzhedvachOnly() tele.MiddlewareFunc {
	lzhedvachID, err := strconv.ParseInt(os.Getenv("LZHEDVACH_ID"), 10, 64)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			if c.Chat().ID == lzhedvachID {
				return nil
			}

			return next(c)
		}
	}
}

package xtelegram

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	TELEGRAM_TOKEN  string
	TELEGRAM_CHATID int
	Bot             *tb.Bot
)

func init() {
  var err error

	TELEGRAM_TOKEN = os.Getenv("TELEGRAM_TOKEN")
	TELEGRAM_CHATID, err = strconv.Atoi(os.Getenv("TELEGRAM_CHATID"))
	if err != nil {
		log.Fatal(err)
	}

	Bot, err = tb.NewBot(tb.Settings{
		Token:  TELEGRAM_TOKEN,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
	}
}

func Send(v interface{}) {
	u := tb.User{
		ID: TELEGRAM_CHATID,
	}

	Bot.Send(&u, fmt.Sprintf("%v", v))
}

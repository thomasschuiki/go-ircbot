package main

import (
	"os"
	"strings"

	"github.com/go-chat-bot/bot/irc"
	_ "github.com/thomasschuiki/go-ircbot/commands/chucknorris"
	_ "github.com/thomasschuiki/go-ircbot/commands/covid"
	_ "github.com/thomasschuiki/go-ircbot/commands/meme"
	_ "github.com/thomasschuiki/go-ircbot/commands/saascat"
)

func main() {
	juicybot := &irc.Config{
		Server:   os.Getenv("IRCSERVER"),
		Channels: strings.Split(os.Getenv("IRCCHANNELS"), ","),
		User:     os.Getenv("IRCUSER"),
		Nick:     os.Getenv("IRCNICK"),
		Debug:    os.Getenv("DEBUG") != ""}

	irc.Run(juicybot)
}

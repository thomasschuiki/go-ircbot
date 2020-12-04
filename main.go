package main

import (
	"os"
	"strings"

	"github.com/go-chat-bot/bot/irc"
	_ "gitlab.com/thomaseitler/juicybotv2/commands/chucknorris"
	_ "gitlab.com/thomaseitler/juicybotv2/commands/meme"
	_ "gitlab.com/thomaseitler/juicybotv2/commands/saascat"
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

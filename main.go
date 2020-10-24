package main

import (
	"os"
	"strings"

	"github.com/go-chat-bot/bot/irc"
	_ "gitlab.com/thomaseitler/juicybotv2/commands/chucknorris/chucknorris"
	_ "gitlab.com/thomaseitler/juicybotv2/commands/saascat/saascat"
)

var (
	ircServer   = "irc.quakenet.org:6667"
	ircChannels = "#uadabotchannel"
	ircUser     = "juicybot"
	ircNick     = "juicybot"
)

func main() {
	juicybot := &irc.Config{
		Server:   ircServer,
		Channels: strings.Split(ircChannels, ","),
		User:     ircUser,
		Nick:     ircNick,
		Debug:    os.Getenv("DEBUG") != ""}

	irc.Run(juicybot)
}

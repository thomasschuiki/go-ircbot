package main

import (
	"log"
	"os"
	"strings"

	"github.com/StalkR/goircbot/bot"
	"github.com/thomasschuiki/go-ircbot/commands/chucknorris"
	"github.com/thomasschuiki/go-ircbot/commands/covid"
	"github.com/thomasschuiki/go-ircbot/commands/meme"
	"github.com/thomasschuiki/go-ircbot/commands/saascat"
)

func main() {
	var host = os.Getenv("IRCSERVER")
	var channels = strings.Split(os.Getenv("IRCCHANNELS"), ",")
	var ident = os.Getenv("IRCUSER")
	var nick = os.Getenv("IRCNICK")
	var ssl = false
	var commandPrefix = "!"

	b, err := bot.NewBotOptions(bot.Host(host), bot.Nick(nick), bot.SSL(ssl), bot.Ident(ident),
		bot.Channels(channels),
		bot.CommandPrefix(commandPrefix))
	if err != nil {
		log.Fatalf("failed to init new bot: %v", err)
	}

  // register plugins

	chucknorris.Register(b)
	covid.Register(b)
	meme.Register(b)
	saascat.Register(b)

  // run the bot
	b.Run()
}

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/go-chat-bot/bot"
	_ "gitlab.com/thomaseitler/juicybotv2/commands/chucknorris"
	_ "gitlab.com/thomaseitler/juicybotv2/commands/meme"
	_ "gitlab.com/thomaseitler/juicybotv2/commands/saascat"
)

func responseHandler(target string, message string, sender *bot.User) {
	if message == "" {
		return
	}
	fmt.Println(fmt.Sprintf("%s: %s", sender.Nick, message))
}

func main() {
	b := bot.New(&bot.Handlers{
		Response: responseHandler,
	},
		&bot.Config{
			Protocol: "debug",
			Server:   "debug",
		},
	)

	fmt.Println("Type a command or !help for available commands...")

	for {
		r := bufio.NewReader(os.Stdin)

		input, err := r.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		b.MessageReceived(
			&bot.ChannelData{
				Protocol:  "debug",
				Server:    "",
				Channel:   "console",
				IsPrivate: true,
			},
			&bot.Message{Text: input},
			&bot.User{ID: "id", RealName: "Debug Console", Nick: "bot", IsBot: false})
	}
}

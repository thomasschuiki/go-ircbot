package chucknorris

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/StalkR/goircbot/bot"
	"github.com/fluffle/goirc/client"
	"github.com/thomasschuiki/go-ircbot/web"
)

const (
	pattern = "(?i)\\b(chuck|norris)\\b"
)

type chuckFact struct {
	Categories []interface{} `json:"categories"`
	CreatedAt  string        `json:"created_at"`
	IconURL    string        `json:"icon_url"`
	ID         string        `json:"id"`
	UpdatedAt  string        `json:"updated_at"`
	URL        string        `json:"url"`
	Value      string        `json:"value"`
}

var (
	baseurl = "https://api.chucknorris.io"
	re      = regexp.MustCompile(pattern)
)

func chucknorris(b bot.Bot, line *client.Line) {
	text := strings.TrimSpace(line.Args[1])
	if re.MatchString(text) {
		url := fmt.Sprintf("%s/jokes/random", baseurl)
		var joke chuckFact

		err := web.MakeAPIRequest(url, nil, nil, &joke)
		if err != nil {
			log.Fatal(err)
		}

		target := line.Args[0]
		b.Privmsg(target, joke.Value)
	}
}

func Register(b bot.Bot) {
	b.HandleFunc("privmsg", func(c *client.Conn, l *client.Line) { chucknorris(b, l) })
}

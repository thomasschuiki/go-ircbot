package chucknorris

import (
	"fmt"
	"regexp"

	"github.com/go-chat-bot/bot"
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

func chucknorris(command *bot.PassiveCmd) (string, error) {
	if re.MatchString(command.Raw) {
		url := fmt.Sprintf("%s/jokes/random", baseurl)
		var joke chuckFact

		err := web.MakeAPIRequest(url, nil, nil, &joke)
		if err != nil {
			return "", err
		}
		return joke.Value, nil
	}

	return "", nil
}

func init() {
	bot.RegisterPassiveCommand(
		"chucknorris", // command
		chucknorris)   // function
}

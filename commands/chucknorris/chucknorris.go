package chucknorris

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/go-chat-bot/bot"
	"gitlab.com/thomaseitler/juicybotv2/web"
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
	baseurl = "https://api.chucknorris.io/"
	re      = regexp.MustCompile(pattern)
)

func chucknorris(command *bot.PassiveCmd) (string, error) {
	if re.MatchString(command.Raw) {
		url := fmt.Sprintf("%s/jokes/random", baseurl)
		var joke chuckFact

		resp, err := web.MakeAPIRequest(url)
		if err != nil {
			return "", err
		}
		parseAPIResponse(resp, &joke)
		return joke.Value, nil
	}

	return "", nil
}

func init() {
	bot.RegisterPassiveCommand(
		"chucknorris", // command
		chucknorris)   // function
}

func parseAPIResponse(r *http.Response, v *chuckFact) error {
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		return fmt.Errorf("json decode error: %v", err)
	}
	return nil
}

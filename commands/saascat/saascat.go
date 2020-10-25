package saascat

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chat-bot/bot"
	"gitlab.com/thomaseitler/juicybotv2/web"
)

type theCatAPIResponse []struct {
	ID         string        `json:"id"`
	URL        string        `json:"url"`
	Breeds     []interface{} `json:"breeds"`
	Categories []interface{} `json:"categories"`
}

var (
	baseurl = "https://api.thecatapi.com/v1"
	apikey  = os.Getenv("SAASCATAPIKEY")
)

func saascat(command *bot.Cmd) (string, error) {
	url := fmt.Sprintf("%s/images/search", baseurl)
	resp, err := web.MakeAPIRequest(url)
	if err != nil {
		return "", err
	}
	var cat theCatAPIResponse
	parseAPIResponse(resp, &cat)
	return cat[0].URL, nil
}

func init() {
	bot.RegisterCommand(
		"cat", // command
		"Returns a random cat pic or gif",
		"!cat",
		saascat) // function
}

func parseAPIResponse(r *http.Response, v *theCatAPIResponse) error {
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		return fmt.Errorf("json decode error: %v", err)
	}
	return nil
}

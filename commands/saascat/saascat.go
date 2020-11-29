package saascat

import (
	"fmt"
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
	header := make(map[string]string)
	header["x-api-key"] = apikey
	queryParams := make(map[string]string)
	// analyze parameters if given
	if len(command.Args) > 0 {
		switch command.Args[0] {
		case "pic":
			queryParams["mime_types"] = "jpg,png"
		case "gif":
			queryParams["mime_types"] = "gif"
		}
	}
	var cat theCatAPIResponse
	err := web.MakeAPIRequest(url, header, queryParams, &cat)
	if err != nil {
		return "", err
	}

	return cat[0].URL, nil
}

func init() {
	bot.RegisterCommand(
		"cat", // command
		"Returns a random cat pic or gif",
		"<pic | gif>",
		saascat) // function
}

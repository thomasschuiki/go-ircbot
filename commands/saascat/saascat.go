package saascat

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
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
	catsearch := fmt.Sprintf("%s/images/search", baseurl)
	header := make(map[string]string)
	header["x-api-key"] = apikey

	// parameter was given
	if len(command.Args) > 0 {
		u, err := url.Parse(catsearch)
		if err != nil {
			log.Fatal(err)
		}
		queryString := u.Query()
		var mimeTypes string
		switch command.Args[0] {
		case "pic":
			mimeTypes = "jpg,png"
		case "gif":
			mimeTypes = "gif"
		}

		queryString.Set("mime_types", mimeTypes)
		u.RawQuery = queryString.Encode()
		catsearch = u.String()
	}

	resp, err := web.MakeAPIRequest(catsearch, header)
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
		"<pic | gif>",
		saascat) // function
}

func parseAPIResponse(r *http.Response, v *theCatAPIResponse) error {
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		return fmt.Errorf("json decode error: %v", err)
	}
	return nil
}

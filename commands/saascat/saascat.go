package saascat

import (
	"fmt"
	"log"
	"os"

	"github.com/StalkR/goircbot/bot"
	"github.com/thomasschuiki/go-ircbot/web"
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

func saascat(e *bot.Event) {
	url := fmt.Sprintf("%s/images/search", baseurl)
	header := make(map[string]string)
	header["x-api-key"] = apikey
	queryParams := make(map[string]string)
	// analyze parameters if given
	if len(e.Args) > 0 {
		switch e.Args {
		case "pic":
			queryParams["mime_types"] = "jpg,png"
		case "gif":
			queryParams["mime_types"] = "gif"
		}
	}
	var cat theCatAPIResponse
	err := web.MakeAPIRequest(url, header, queryParams, &cat)
	if err != nil {
		log.Fatal(err)
	}

  returnString := cat[0].URL
	e.Bot.Privmsg(e.Target, returnString)
}

// Register registers the plugin with a bot.
func Register(b bot.Bot) {
	b.Commands().Add("cat", bot.Command{
		Help:    "Returns a random cat pic or gif. Supply 'pic' or 'gif' to get a specific image.",
		Handler: saascat,
		Pub:     true,
		Priv:    true,
		Hidden:  false})
}

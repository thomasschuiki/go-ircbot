package meme

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/StalkR/goircbot/bot"
	"github.com/thomasschuiki/go-ircbot/web"
)

var (
	baseurl = "https://knowyourmeme.com"
)

func meme(e *bot.Event) {
	var url, title, href string
	queryParams := make(map[string]string)
	isSearch := false
	// analyze parameters if given
	if len(e.Args) > 0 {
		url = fmt.Sprintf("%s/search", baseurl)
		queryParams["q"] = e.Args
		isSearch = true
	} else {
		url = fmt.Sprintf("%s/random", baseurl)
	}
	response, err := web.GetWebpage(url, queryParams)
	if err != nil {
		log.Fatal(err)
	}
	memedoc, err := goquery.NewDocumentFromReader(response)
	if err != nil {
		log.Fatal(err)
	}
	var sb strings.Builder
	if !isSearch {
		// Random Meme
		link := memedoc.Find("section.info > h1:nth-child(1) > a:nth-child(1)").First()

		// For each item found, get the title and url
		title = link.Text()
		href, _ = link.Attr("href")
		sb.WriteString(fmt.Sprintf("%s %s\n", title, fmt.Sprintf("%s%s", baseurl, href)))
	} else {
		// Get first 5 results of search
		memedoc.Find(".entry-grid-body > tr:nth-child(1) td").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the title and url
			link := s.Find("h2 a")
			title = link.Text()
			href, _ = link.Attr("href")
			sb.WriteString(fmt.Sprintf("%d. %s %s\n", i, title, fmt.Sprintf("%s%s", baseurl, href)))
		})
	}

	returnString := sb.String()
	e.Bot.Privmsg(e.Target, returnString)
}

// Register registers the plugin with a bot.
func Register(b bot.Bot) {
	b.Commands().Add("meme", bot.Command{
		Help:    "Returns a random meme or searches for one on knowyourmeme.org. Use 'random' or the name of a meme as an argument.",
		Handler: meme,
		Pub:     true,
		Priv:    true,
		Hidden:  false})
}

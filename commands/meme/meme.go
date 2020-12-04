package meme

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-chat-bot/bot"
	"gitlab.com/thomaseitler/juicybotv2/web"
)

var (
	baseurl = "https://knowyourmeme.com"
)

func meme(command *bot.Cmd) (string, error) {
	var url, title, href string
	queryParams := make(map[string]string)
	isSearch := false
	// analyze parameters if given
	if len(command.Args) > 0 {
		url = fmt.Sprintf("%s/search", baseurl)
		queryParams["q"] = command.Args[0]
		isSearch = true
	} else {
		url = fmt.Sprintf("%s/random", baseurl)
	}
	response, err := web.GetWebpage(url, queryParams)
	if err != nil {
		return "", err
	}
	memedoc, err := goquery.NewDocumentFromReader(response)
	if err != nil {
		return "", err
	}
	var sb strings.Builder
	if !isSearch {
		// Random Meme
		link := memedoc.Find("section.info > h1:nth-child(1) > a:nth-child(1)").First()

		// For each item found, get the title and url
		// link := s.Find("h2 a")
		title = link.Text()
		href, _ = link.Attr("href")

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

	botmsg := sb.String()
	// fmt.Println(botmsg)
	return botmsg, nil
}

func init() {
	bot.RegisterCommand(
		"meme", // command
		"Returns a random meme or searches for one on knowyourmeme.org",
		"<random | memename>",
		meme) // function
}

package covid

import (
	"fmt"

	"github.com/go-chat-bot/bot"
	"github.com/thomasschuiki/go-ircbot/web"
)

type covidResponse struct {
	Updated     int64  `json:"updated"`
	Country     string `json:"country"`
	CountryInfo struct {
		ID   int     `json:"_id"`
		Iso2 string  `json:"iso2"`
		Iso3 string  `json:"iso3"`
		Lat  float64 `json:"lat"`
		Long float64 `json:"long"`
		Flag string  `json:"flag"`
	} `json:"countryInfo"`
	Cases                  int     `json:"cases"`
	TodayCases             int     `json:"todayCases"`
	Deaths                 int     `json:"deaths"`
	TodayDeaths            int     `json:"todayDeaths"`
	Recovered              int     `json:"recovered"`
	TodayRecovered         int     `json:"todayRecovered"`
	Active                 int     `json:"active"`
	Critical               int     `json:"critical"`
	CasesPerOneMillion     int     `json:"casesPerOneMillion"`
	DeathsPerOneMillion    int     `json:"deathsPerOneMillion"`
	Tests                  int     `json:"tests"`
	TestsPerOneMillion     int     `json:"testsPerOneMillion"`
	Population             int     `json:"population"`
	Continent              string  `json:"continent"`
	OneCasePerPeople       int     `json:"oneCasePerPeople"`
	OneDeathPerPeople      int     `json:"oneDeathPerPeople"`
	OneTestPerPeople       int     `json:"oneTestPerPeople"`
	ActivePerOneMillion    float64 `json:"activePerOneMillion"`
	RecoveredPerOneMillion float64 `json:"recoveredPerOneMillion"`
	CriticalPerOneMillion  float64 `json:"criticalPerOneMillion"`
}

var (
	baseurl = "https://disease.sh/v3"
)

func covid(command *bot.Cmd) (string, error) {
	url := fmt.Sprintf("%s/covid-19/countries", baseurl)
	header := make(map[string]string)
	queryParams := make(map[string]string)
	// analyze parameters if given
	if len(command.Args) > 0 {
		url = fmt.Sprintf("%s/%s", url, command.Args[0])
		var cR covidResponse
		err := web.MakeAPIRequest(url, header, queryParams, &cR)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("Todays Cases: %d", cR.TodayCases), nil
	}
	return "couldn't find any information", nil
}

func init() {
	bot.RegisterCommand(
		"covid", // command
		"Returns statistics about CVOID via disease.sh",
		"<country>",
		covid) // function
}

package covid

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/StalkR/goircbot/bot"
	"github.com/thomasschuiki/go-ircbot/web"
)

type wmCountriesCountryResponse struct {
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

type jhucsseHistoricalCountryResponse struct {
	Country  string `json:"country"`
	Timeline struct {
		Cases     []jhucsseStat `json:"cases"`
		Deaths    []jhucsseStat `json:"deaths"`
		Recovered []jhucsseStat `json:"recovered"`
	} `json:"timeline"`
}
type jhucsseStat struct {
	Date  time.Time
	Count float64
}

func (j *jhucsseHistoricalCountryResponse) UnmarshalJSON(data []byte) error {

	var v map[string]interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	const (
		layoutUS = "1/2/06"
	)

	j.Country, _ = v["country"].(string)

	t := v["timeline"].(map[string]interface{})
	cases := t["cases"].(map[string]interface{})
	for d, c := range cases {
		parsedTime, err := time.Parse(layoutUS, d)
		if err != nil {
			return fmt.Errorf("error parsing time: %+v", err)
		}
		s := jhucsseStat{Count: c.(float64), Date: parsedTime}
		j.Timeline.Cases = append(j.Timeline.Cases, s)
	}

	return nil
}

var (
	baseurl = "https://disease.sh/v3"
)

func covid(e *bot.Event) {
	countryurl := fmt.Sprintf("%s/covid-19/countries", baseurl)
	histurl := fmt.Sprintf("%s/covid-19/historical", baseurl)
	header := make(map[string]string)
	queryParams := make(map[string]string)
	lastdays := 8
	args := strings.Split(e.Args, " ")
	// analyze parameters if given
	if len(args) > 0 {
		countryurl = fmt.Sprintf("%s/%s", countryurl, args[0])
		histurl = fmt.Sprintf("%s/%s", histurl, args[0])

		var cRYesterday wmCountriesCountryResponse
		var cRTwoDaysAgo wmCountriesCountryResponse
		var casesToday wmCountriesCountryResponse
		err := web.MakeAPIRequest(countryurl, header, nil, &casesToday)
		if err != nil {
			log.Fatal(err)
		}
		err = web.MakeAPIRequest(countryurl, header, map[string]string{"yesterday": "true"}, &cRYesterday)
		if err != nil {
			log.Fatal(err)
		}
		err = web.MakeAPIRequest(countryurl, header, map[string]string{"twoDaysAgo": "true"}, &cRTwoDaysAgo)
		if err != nil {
			log.Fatal(err)
		}
		var hist7 jhucsseHistoricalCountryResponse
		queryParams["lastdays"] = fmt.Sprintf("%d", lastdays)
		err = web.MakeAPIRequest(histurl, header, queryParams, &hist7)
		if err != nil {
			log.Fatal(err)
		}
		incidence_rate := calculateIncidence(hist7.Timeline.Cases, casesToday.Population, lastdays, casesToday.TodayCases)
		var strToday string
		strChange := ""
		if casesToday.TodayCases > 0 {
			strToday = fmt.Sprintf("Cases today: %d", casesToday.TodayCases)
			ChangeSinceYesterday := percentageChange(cRYesterday.TodayCases, casesToday.TodayCases)
			ChangeSinceTwoDaysAgo := percentageChange(cRTwoDaysAgo.TodayCases, casesToday.TodayCases)
			strChange = fmt.Sprintf("\nThat is %3.2f%% since Yesterday and %3.2f%% since 2-days ago.", ChangeSinceYesterday, ChangeSinceTwoDaysAgo)
		} else {
			strToday = "Cases today: no data"
		}

		returnString := fmt.Sprintf("%s, Cases yesterday: %d, Cases 2-days ago: %d.%s\n7-day incidency rate is ~%3.3f", strToday, cRYesterday.TodayCases, cRTwoDaysAgo.TodayCases, strChange, incidence_rate)
		e.Bot.Privmsg(e.Target, returnString)
    return
	}
	returnString := "Please provide a country name, iso2, iso3, or country ID code. e.g.: AT, Austria "
	e.Bot.Privmsg(e.Target, returnString)
  return
}

func percentageChange(old, new int) float64 {
	diff := float64(new - old)
	return (diff / float64(old)) * 100
}

func calculateIncidence(histCases []jhucsseStat, population, lastdays, casesToday int) float64 {
	// sort slice by date
	sort.Slice(histCases, func(i, j int) bool {
		return histCases[i].Date.Before(histCases[j].Date)
	})
	// calculate daily cases
	// var cases_in_period_by_day []float64
	cases_in_period_by_day_sum := 0.0
	for i := range histCases {
		if i >= lastdays-1 {
			break
		}
		diff := float64(histCases[i+1].Count) - float64(histCases[i].Count)
		cases_in_period_by_day_sum += diff
	}
	// cases_in_period_by_day_sum += float64(casesToday)
	// calculate incidence
	return cases_in_period_by_day_sum / float64(population) * 100000.0
}

// Register registers the plugin with a bot.
func Register(b bot.Bot) {
	b.Commands().Add("covid", bot.Command{
		Help:    "Returns statistics about COVID 19 via disease.sh. Provide country as argument.",
		Handler: covid,
		Pub:     true,
		Priv:    true,
		Hidden:  false})
}

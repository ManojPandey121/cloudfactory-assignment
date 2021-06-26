package covid

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UserAgent interface {
	Get(url string) (*http.Response, error)
}

type CountryInfo struct {
	Id   int    `json:"_id"`
	Flag string `json:"flag"`
}

var DefaultClient = &http.Client{}

type CountryRecord struct {
	Updated                int         `json:"updated"`
	Country                string      `json:"country"`
	CountryInfo            CountryInfo `json:"countryInfo"`
	Cases                  int         `json:"cases"`
	TodayCases             int         `json:"todayCases"`
	Deaths                 int         `json:"deaths"`
	TodayDeaths            int         `json:"todayDeaths"`
	Recovered              int         `json:"recovered"`
	TodayRecovered         int         `json:"todayRecovered"`
	Active                 int         `json:"active"`
	Critical               int         `json:"critical"`
	CasesPerOneMillion     float32     `json:"casesPerOneMillion"`
	DeathsPerOneMillion    float32     `json:"deathsPerOneMillion"`
	Tests                  int         `json:"tests"`
	TestsPerOneMillion     float32     `json:"testsPerOneMillion"`
	Population             int         `json:"population"`
	Continent              string      `json:"continent"`
	OneCasePerPeople       float32     `json:"oneCasePerPeople"`
	OneDeathPerPeople      float32     `json:"oneDeathPerPeople"`
	OneTestPerPeople       float32     `json:"oneTestPerPeople"`
	Undefined              float32     `json:"undefined"`
	ActivePerOneMillion    float32     `json:"activePerOneMillion"`
	RecoveredPerOneMillion float32     `json:"recoveredPerOneMillion"`
	CriticalPerOneMillion  float32     `json:"criticalPerOneMillion"`
}

const url = "https://corona.lmao.ninja/v2/countries?yesterday=&sort="

func GetCoviddata(countryCase string) {
	country := []CountryRecord{}
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(data, &country)
	if err != nil {
		fmt.Printf("Error unmarshilling http data %+v", err)
		return
	}

	for _, covidRecord := range country {
		if countryCase != "" && covidRecord.Country == countryCase {
			fmt.Printf("Covid records for %s \n", covidRecord.Country)
			fmt.Printf("Total cases: %d\t Cases Found Today: %d\n ", covidRecord.Cases, covidRecord.TodayCases)
			fmt.Printf("Total Recovered: %d\t  Total Deaths: %d\n", covidRecord.Recovered, covidRecord.Deaths)
			fmt.Printf("======>\n")
			break
		} else if countryCase == "" || countryCase == "all" {
			fmt.Printf("Covid records for %s \n", covidRecord.Country)
			fmt.Printf("Total cases: %d\t Cases Found Today: %d\n ", covidRecord.Cases, covidRecord.TodayCases)
			fmt.Printf("Total Recovered: %d\t  Total Deaths: %d\n", covidRecord.Recovered, covidRecord.Deaths)
			fmt.Printf("======>\n")
		}
	}
}

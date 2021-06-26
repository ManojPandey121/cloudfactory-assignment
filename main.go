package main

import (
	"net/http"
)

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

func main() {
}

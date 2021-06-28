package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	ALL_COUNTRY_URL = "https://corona.lmao.ninja/v2/countries?yesterday=&sort="
	ONE_COUNTRY_URL = "https://corona.lmao.ninja/v2/countries/"
)

func GetCoviddata(countryCase string) {
	allCountries := []CountryRecord{}
	singleCountry := CountryRecord{}
	if countryCase == "" || countryCase == "all" {
		response, err := http.Get(ALL_COUNTRY_URL)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
			return
		}
		defer response.Body.Close()
		data, _ := ioutil.ReadAll(response.Body)
		err = json.Unmarshal(data, &allCountries)
		if err != nil {
			fmt.Printf("Error unmarshilling http data %+v", err)
			return
		}
		for _, covidRecord := range allCountries {
			PrettyPrint(covidRecord)
		}
	} else {
		response, err := http.Get(ONE_COUNTRY_URL + countryCase)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
			return
		}
		defer response.Body.Close()
		data, _ := ioutil.ReadAll(response.Body)
		err = json.Unmarshal(data, &singleCountry)
		if err != nil {
			fmt.Printf("Error unmarshilling http data %+v", err)
			return
		}
		if singleCountry.Country != "" {
			PrettyPrint(singleCountry)
		} else {
			fmt.Printf("Sorry, data not found for %s country \n", countryCase)
		}
	}

}

func PrettyPrint(countryRecord CountryRecord) {
	fmt.Printf("------Covid records for %s ----- \n", countryRecord.Country)
	fmt.Printf("Total cases: %d\t Cases Found Today: %d\n ", countryRecord.Cases, countryRecord.TodayCases)
	fmt.Printf("Total Recovered: %d\t  Total Deaths: %d\n", countryRecord.Recovered, countryRecord.Deaths)
	fmt.Printf("======\n")
}

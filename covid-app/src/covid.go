package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL = "https://corona.lmao.ninja/v2/countries?yesterday=&sort="

func GetCoviddata(countryCase string) {
	country := []CountryRecord{}
	response, err := http.Get(URL)
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
		PrettyPrint(covidRecord)
	}
}

func PrettyPrint(countryRecord CountryRecord) {
	fmt.Printf("Covid records for %s \n", countryRecord.Country)
	fmt.Printf("Total cases: %d\t Cases Found Today: %d\n ", countryRecord.Cases, countryRecord.TodayCases)
	fmt.Printf("Total Recovered: %d\t  Total Deaths: %d\n", countryRecord.Recovered, countryRecord.Deaths)
	fmt.Printf("======>\n")
}

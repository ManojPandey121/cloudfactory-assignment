package main

import (
	"covid/covid"
	"fmt"
)

func main() {
	fmt.Println("Starting the application...")
	var country string
	fmt.Print("write country name to get records (Type 'all' to get records of all countries):")
	fmt.Scanf("%s\n", &country)
	covid.GetCoviddata(country)
	fmt.Println("Terminating the application...")
}

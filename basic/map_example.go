package main

import (
	"fmt"
)

// var map_variable map[key_data_type]value_data_type
// map_variable := make(map[key_data_type]value_data_type)
// 没有初始化map时，会创建nil map

func main() {
	var countryCapital map[string]string
	countryCapital = make(map[string]string)

	countryCapital["France"] = "Paris"
	countryCapital["China"] = "Beijing"
	countryCapital["American"] = "纽约"

	for country := range countryCapital {
		fmt.Println(country, "Capital is", countryCapital[country])
	}

	capital, ret := countryCapital["Japan"]
	if ret {
		fmt.Println("Japan Capital is", capital)
	} else {
		fmt.Println("Japan Capital not found!")
	}

	// delete
	delete(countryCapital, "American")
	fmt.Println("Delete American")
	for country := range countryCapital {
		fmt.Println(country, "=>", countryCapital[country])
	}
}

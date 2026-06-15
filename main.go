package main

import (
	"fmt"
)

func main() {
	// variables
	var name string = "Rendra"
	fmt.Printf("This is my name %s\n", name)

	age := 41
	fmt.Printf("This is my age %d\n", age)

	var city string
	city = "Bogor"
	fmt.Printf("this is my city %s\n", city)

	var country, continent string = "IDN", "Asia"
	fmt.Printf("This is my country %s and my continent %s\n", country, continent)

	var (
		isEmployed bool   = true
		salary     int    = 50000
		position   string = "developer"
	)

	fmt.Printf("is employed : %t, this is my salary %d and my position is %s\n", isEmployed, salary, position)

	// Zero values
	var defaultInt int
	var defaultFloat float64
	var defaultString string
	var defaultBool bool

	fmt.Printf("int: %d, float: %f, string: %s, bool: %t\n", defaultInt, defaultFloat, defaultString, defaultBool)

}

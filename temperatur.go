package main

import "fmt"

func main() {
	var temp, convert float64
	var unit string
	fmt.Scan(&temp, &unit)
	if unit == "Celcius" || unit == "celcius" {
		convert = 32 + (9*temp)/5
		fmt.Printf("Suhu dalam Fahrenheit adalah %.0f", convert)
	} else if unit == "Fahrenheit" || unit == "fahrenheit" {
		convert = (temp + (-32)) * 5 / 9
		fmt.Printf("Suhu dalam Celcius adalah %.0f", convert)
	} else {
		fmt.Print("Sintax Error :(")
	}
}

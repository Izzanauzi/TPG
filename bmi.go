package main

import "fmt"

func main() {
	var tb, bb int
	var bmi float64
	fmt.Scan(&tb, &bb)
	bmi = float64(bb / ((tb / 100) ^ 2))
	if bmi > 22.9 {
		fmt.Println("Berat Badan Berlebih")
	} else if bmi <= 22.9 && bmi >= 18.5 {
		fmt.Println("Berat Badan Normal")
	} else {
		fmt.Println("Berat Badan Kurang")
	}
}

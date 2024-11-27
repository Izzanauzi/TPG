package main

import "fmt"

func main() {
	var a float64
	fmt.Scan(&a)
	for i := 0; i <= 10; i++ {
		a = a * 1.05
		fmt.Printf("Tahun ke-%d : Rp.%.2f\n", i, a)
	}
}

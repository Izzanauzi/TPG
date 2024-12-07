package main

import "fmt"

func main() {
	var plat, hasil, final int
	fmt.Scan(&plat)
	hasil = 0
	final = 0
	for i := 0; i < 5; i++ {
		hasil = hasil + (plat % 10)
		plat = plat / 10
	}
	for i := 0; i < 2; i++ {
		final = final + (hasil % 10)
		hasil = hasil / 10
	}
	fmt.Println(final)
}

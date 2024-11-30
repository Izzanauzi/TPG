package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n)
	x = 0
	for n > 0 {
		n = n / 10
		x = x + 1
	}
	fmt.Printf("%d digit", x)
}

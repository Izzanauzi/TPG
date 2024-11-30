package main

import "fmt"

func main() {
	var n, x, y int
	var z bool
	fmt.Scan(&n)
	z = true
	for n > 0 {
		x = n % 10
		y = n / 10 % 10
		z = z && (x-y == 1 || y-x == 1)
		n = n / 10
	}
	fmt.Println(z)
}

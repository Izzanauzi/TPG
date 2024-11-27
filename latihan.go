package main

import "fmt"

func main() {
	var i, x, m, n int
	fmt.Scan(&x, &m, &n)

	for i = (x - m); i <= (x - n); i++ {
		fmt.Print(i, " ")
	}
}

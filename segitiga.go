package main

import "fmt"

func main() {
	var n int
	var a, b, c float64
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a, &b)
		c = (a * b) / 2
		fmt.Println(c)
	}
}

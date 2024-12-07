package main

import "fmt"

func main() {
	var n int
	var v, y, g float64
	fmt.Scan(&n, &v)
	g = 9.8
	for i := 0; i <= n; i++ {
		fmt.Println(y)
		y = float64(i) + v + ((0.5 * g) * (0.5 * g))
	}
}

package main

import "fmt"

func main() {
	var a, a1, a2, a3 int
	fmt.Scan(&a)
	a1 = a / 100
	a2 = a % 100 / 10
	a3 = a % 10
	fmt.Print(a1 * 11)
	fmt.Print(a2 * 11)
	fmt.Print(a3 * 11)
}

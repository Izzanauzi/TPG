package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b)
	for i := b; i <= a; i++ {
		fmt.Print(i)
		c += i
	}
	fmt.Print("\n", c)
}

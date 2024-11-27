package main

import "fmt"

func main() {
	var a int
	var b bool
	fmt.Scan(&a)
	b = a%5 == 0 && a%4 == 0 && a%100 != 0
	fmt.Print(b)
}

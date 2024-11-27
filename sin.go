package main

import "fmt"

func main() {
	var a, b int
	var c bool
	fmt.Print("Bilangan x dan N : ")
	fmt.Scan(&a, &b)
	c = a%b == 0
	fmt.Printf("%d kelipatan %d ? %t", a, b, c)
}

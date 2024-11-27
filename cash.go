package main

import "fmt"

func main() {
	var a int
	var b bool
	fmt.Scan(&a)
	b = a/1000 == a%10
	fmt.Print(b)
}

package main

import (
	"fmt"
)

func main() {
	var a, b, c float64
	var segitiga bool
	fmt.Scan(&a, &b, &c)
	segitiga = (a > 0 && b > 0 && c > 0) && (a+b > c && b+c > a && c+a > b)
	fmt.Print(a, ",", b, ", dan ", c, " segitiga ? ", segitiga)
}

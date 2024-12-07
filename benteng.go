package main

import "fmt"

func main() {
	var n, x, y, nn, xx, yy int
	fmt.Scan(&n, &x, &y)
	xx = 1
	yy = 1
	for a := false; !a; {
		if xx != y {
			fmt.Printf("(%d , %d)\n", x, xx)
		}
		if yy != x {
			fmt.Printf("(%d , %d)\n", yy, y)
		}
		xx++
		yy++
		nn++
		a = nn == n
	}
}

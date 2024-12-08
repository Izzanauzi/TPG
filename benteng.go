package main

import "fmt"

func main() {
	var n, x, y, nn, xx, yy int
	fmt.Scan(&n, &x, &y)
	xx = 1
	yy = 1
	for a := false; !a; {
		if yy != y {
			fmt.Printf("(%d , %d)\n", x, yy)
		}
		if xx != x {
			fmt.Printf("(%d , %d)\n", xx, y)
		}
		xx++
		yy++
		nn++
		a = nn == n
	}
}

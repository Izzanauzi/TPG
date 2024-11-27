package main

import "fmt"

func main() {
	var need, x int
	fmt.Scan(&need, &x)
	if need <= 1000 {
		fmt.Print(1*x, " Karung")
	} else if need > 1000 && need <= 2000 {
		fmt.Print(2*x, " Karung")
	} else if need > 2000 && need <= 3000 {
		fmt.Print(3*x, " Karung")
	}
}

package main

import "fmt"

func main() {
	var adik, kakak int
	var tebak bool
	fmt.Scan(&adik, &kakak)
	tebak = adik == kakak || (adik+1) == kakak || (adik-1) == kakak || (adik+5) == kakak || (adik-5) == kakak
	fmt.Print(tebak)
}

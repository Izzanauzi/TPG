package main

import "fmt"

func main() {
	var anime, total, n, min, max int
	var avg float64
	min = 100
	max = 0
	for power := false; !power; {
		fmt.Scan(&anime)
		if anime > 0 && anime < 100 {
			n++
			total = total + anime
			avg = float64(total / n)
			if min > anime {
				min = anime
			}
			if max < anime {
				max = anime
			}
		}
		power = anime == 100
	}
	if n > 0 {
		fmt.Printf("%d %.2f %d %d", n, avg, min, max)
	}
}

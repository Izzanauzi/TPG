package main

import "fmt"

func main() {
	var n, score, total int
	n = 0
	total = 0
	for condition := false; !condition; {
		fmt.Scan(&score)
		total = total + score
		n = n + 1
		condition = total >= 30
	}
	fmt.Println(total, n)
}

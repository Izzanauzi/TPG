package main

import "fmt"

func main() {
	var m, n, x int
	var nilai bool
	nilai = true
	fmt.Scan(&m, &n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&x)
		nilai = nilai && (x == m*i)
	}
	fmt.Println(nilai)
}

package main

import "fmt"

func main() {
	var transaksi float64
	fmt.Scanf("%f", &transaksi)
	if transaksi >= 500 && transaksi <= 1000 {
		fmt.Println(transaksi - (transaksi * 0.15))
	} else if transaksi > 1000 {
		fmt.Println(transaksi - (transaksi * 0.20))
	} else {
		fmt.Println(transaksi - (transaksi * 0.05))
	}
}

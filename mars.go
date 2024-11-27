package main

import "fmt"

func main() {
	var detikInput, detikSisa, jam, menit, detik int
	fmt.Print("Input dalam satuan detik : ")
	fmt.Scan(&detikInput)
	jam = detikInput / 4500
	detikSisa = detikInput % 4500
	menit = detikSisa / 75
	detik = detikSisa % 75
	fmt.Println("Maka hasil konversinya adalah ", jam, " jam ", menit, " menit dan ", detik, "detik di mars")
}

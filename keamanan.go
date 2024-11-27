package main

import "fmt"

func main() {
	var kartu, surat bool
	fmt.Scan(&kartu, &surat)
	if kartu == true || surat == true {
		fmt.Print("Diizinkan Masuk")
	} else {
		fmt.Print("Tidak Diizinkan Masuk")
	}
}

package main

import "fmt"

func main() {
	var pin, nominal, n int
	for validation := false; !validation; {
		fmt.Print("PIN: ")
		fmt.Scan(&pin)
		n++
		validation = pin == 123987 || n >= 3
	}
	if n >= 3 {
		fmt.Println("PIN terblokir, silahkan kunjungi kantor cabang terdekar")
	} else {
		fmt.Print("Nominal uang: ")
		fmt.Scan(&nominal)
		fmt.Printf("Uang Rp%d sudah bisa diambil, terima kasih.", nominal)
	}
}

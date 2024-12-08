package main

import "fmt"

func main() {
	var pin, nominal, n int
	n = 0
	for validation := false; !validation; {
		fmt.Print("PIN: ")
		fmt.Scan(&pin)
		n++
		validation = pin == 123987 || n == 3
		if validation == false {
			fmt.Println("PIN Salah")
		}
	}
	if pin != 123987 {
		fmt.Println("PIN terblokir, silahkan kunjungi kantor cabang terdekar")
	} else {
		fmt.Print("Nominal uang: ")
		fmt.Scan(&nominal)
		fmt.Printf("Uang Rp%d sudah bisa diambil, terima kasih.", nominal)
	}
}

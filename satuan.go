package main

import "fmt"

func main() {
	var nominal int
	fmt.Scan(&nominal)
	switch {
	case nominal >= 0 && nominal <= 9:
		fmt.Println("Satuan")
	case nominal >= 10 && nominal <= 99:
		fmt.Println("Puluhan")
	case nominal >= 100 && nominal <= 999:
		fmt.Println("Ratusan")
	case nominal >= 1000 && nominal <= 9999:
		fmt.Println("Ribuan")
	case nominal >= 10000 && nominal <= 99999:
		fmt.Println("Puluhan Ribu")
	case nominal >= 100000 && nominal <= 999999:
		fmt.Println("Ratusan Ribu")
	}

}

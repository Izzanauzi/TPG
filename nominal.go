package main

import "fmt"

func main() {
	var nominal int
	fmt.Scanf("%d", &nominal)
	if nominal/100000 >= 1 {
		fmt.Println("Ratusan Ribu")
	} else if nominal/10000 >= 1 {
		fmt.Println("Puluhan Ribu")
	} else if nominal/1000 >= 1 {
		fmt.Println("Ribuan")
	} else if nominal/100 >= 1 {
		fmt.Println("Ratusan")
	} else if nominal/10 >= 1 {
		fmt.Println("Puluhan")
	} else {
		fmt.Println("Satuan")
	}
}

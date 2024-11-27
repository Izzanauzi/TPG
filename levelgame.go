package main

import "fmt"

func main() {
	var level int
	fmt.Scan(&level)
	switch {
	case level >= 1 && level <= 10:
		fmt.Println("Pemula")
	case level >= 11 && level <= 20:
		fmt.Println("Menengah")
	case level >= 21 && level <= 30:
		fmt.Println("Ahli")
	case level > 30:
		fmt.Println("Master")
	default:
		fmt.Println("Level tidak valid")
	}
}

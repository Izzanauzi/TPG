package main

import "fmt"

func main() {
	var a, b, hasil float64
	var operator rune
	fmt.Scanf("%f %c %f", &a, &operator, &b)
	switch operator {
	case '+':
		hasil = a + b
	case '-':
		hasil = a - b
	case '*':
		hasil = a * b
	case '/':
		hasil = a / b
	}
	fmt.Printf("%.3f", hasil)
}

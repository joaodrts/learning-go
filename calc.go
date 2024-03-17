package main

import "fmt"

func main() {

	// Outher example: n1 := 2
	var n1 int
	var n2 int

	n1 = 2
	n2 = 20

	fmt.Println("O resultado da soma é:", Sum(n1, n2))
	fmt.Println("O resultado da subtração é:", Subtract(n1, n2))
	fmt.Println("O resultado da multiplicação é:", Multiplication(n1, n2))
	fmt.Println("O resultado da divisão é:", Division(n1, n2))
}

func Sum(n1 int, n2 int) int {
	return n1 + n2
}

func Subtract(n1 int, n2 int) int {
	return n1 - n2
}

func Multiplication(n1 int, n2 int) int {
	return n1 * n2
}

func Division(n1 int, n2 int) int {

	if n2 <= 0 {
		return 0
	}

	return n1 / n2
}

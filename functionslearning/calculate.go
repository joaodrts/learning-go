package functionslearning

import "fmt"

func Calculate(n1 int, n2 int) {

	fmt.Println("A soma é", sum(n1, n2))
	fmt.Println("A subtração é", subtract(n1, n2))
	fmt.Println("A multiplicação é", multiplication(n1, n2))
	fmt.Println("A divisão é", division(n1, n2))

}

func sum(n1 int, n2 int) int {
	return n1 + n2
}

func subtract(n1 int, n2 int) int {
	return n1 - n2
}

func multiplication(n1 int, n2 int) int {
	return n1 * n2
}

func division(n1 int, n2 int) int {

	if n2 <= 0 {
		return 0
	}

	return n1 / n2

}

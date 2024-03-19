package main

import (
	"fmt"
)

type Carro struct {
	Nome  string
	Marca string
	Ano   int
}

func main() {
	var newCarro Carro = CreateCar("Gol", "Volks", 2000)
	fmt.Println(newCarro)
}

func CreateCar(nome string, marca string, ano int) Carro {

	return Carro{Ano: 2000, Marca: "Volks", Nome: "Gol"}

}

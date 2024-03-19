package main

import (
	f "functionslearning"
)

func main() {

	f.Create("Gol g3", "Volkswagem", 2004)
	f.Calculate(10, 20)

	persons := []f.Person{
		{Name: "joao", Age: 21},
		{Name: "duarte", Age: 98},
		{Name: "Silva", Age: 30},
	}

	f.CratePerson(persons)

	person := f.Person{Name: "João Duarte", Age: 21}
	products := []f.Product{
		{Name: "Product 01", Value: 10.50},
		{Name: "Product 02", Value: 119},
		{Name: "Product 03", Value: 0.99},
	}

	f.CreateSale(person, products)
}

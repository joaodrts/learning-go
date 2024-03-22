package main

import (
	"fmt"
	f "learning-go/functions"
)

func main() {

	f.Create("Gol g3", "Volkswagem", 2004)
	f.Calculate(10, 20)

	person := f.Person{Name: "João Duarte", Age: 21}
	products := []f.Product{
		{Name: "Product 01", Value: 10.50},
		{Name: "Product 02", Value: 119},
		{Name: "Product 03", Value: 0.99},
	}

	f.CreateSale(person, products)

	fmt.Println("Inicio do CRUD...")

	fmt.Println("Inicio do CRUD - GetAll...")

	result := f.GelAll()
	fmt.Println(result)

	fmt.Println("Inicio do Add...")
	f.Add(f.Client{Fullname: "Joao", Age: 21})
	f.Add(f.Client{Fullname: "Joao G", Age: 21})

	fmt.Println("Inicio do GetAll...")
	resultGetAll := f.GelAll()
	fmt.Println(resultGetAll)

	fmt.Println("Inicio do GetById...")
	resultGetById := f.GetById(1)
	fmt.Println(resultGetById)

	fmt.Println("Inicio do Delete...")
	f.Delete(1)
	resultGetAll2 := f.GelAll()

	fmt.Println(resultGetAll2)

}

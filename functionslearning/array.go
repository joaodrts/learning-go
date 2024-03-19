package functionslearning

import "fmt"

type Sale struct {
	id       int
	client   Person
	products []Product
	value    float64
}

type Person struct {
	Name string
	Age  int
}

type Product struct {
	Name  string
	Value float64
}

func CratePerson(persons []Person) {

	list := make([]Person, len(persons))

	for i := 0; i < len(persons); i++ {
		list[i] = persons[i]
	}

	fmt.Println(list)
	fmt.Println("Tamanho do array: ", len(list))
}

func CreateSale(person Person, products []Product) {

	var sale Sale = Sale{id: 1, client: person, products: products, value: sumProduct(products)}

	fmt.Println("Client: ", sale.client.Name)
	fmt.Println("Products: ", sale.products)
	fmt.Println("Sale Value: ", sale.value)

}

func sumProduct(products []Product) float64 {

	var valueProducts float64

	for i := 0; i < len(products); i++ {
		valueProducts = valueProducts + products[i].Value
	}

	return valueProducts

}

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

func CreateSale(person Person, products []Product) {

	// or if result := validationsForSale(person, products); result == true {
	if validationsForSale(person, products) {

		var sale Sale = Sale{id: 1, client: person, products: products, value: sumProduct(products)}

		fmt.Println("Client: ", sale.client.Name)
		fmt.Println("Products: ", sale.products)
		fmt.Println("Sale Value: ", sale.value)
	}

}

func sumProduct(products []Product) float64 {

	var valueProducts float64

	for i := 0; i < len(products); i++ {
		valueProducts += products[i].Value
	}

	return valueProducts

}

func validationsForSale(person Person, products []Product) bool {

	if person.Name == "" || person.Age == 0 {
		fmt.Println("Report a customer.")
		return false
	} else if len(products) == 0 {
		fmt.Println("inform at least one product")
		return false
	}

	return true
}

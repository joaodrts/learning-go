package functionslearning

import "fmt"

type carro struct {
	nome  string
	marca string
	ano   int
}

func Create(nome string, marca string, ano int) {

	var newCar carro = carro{nome: nome, marca: marca, ano: ano}
	fmt.Println(newCar)

}

package functionslearning

type Client struct {
	Id       int
	Fullname string
	Age      int
}

var listClient []Client

func Add(client Client) {
	client.Id = len(listClient) + 1
	listClient = append(listClient, client)
}

func GelAll() []Client {
	return listClient
}

func Edit(id int, client Client) {

	for x := 0; x < len(listClient); x++ {
		if listClient[x].Id == id {
			listClient[x].Fullname = client.Fullname
			listClient[x].Age = client.Age
			break
		}
	}

}

func Delete(id int) {

	var newList []Client

	for x := 0; x < len(listClient); x++ {
		if listClient[x].Id != id {
			newList = append(newList, listClient[x])
		}
	}

	listClient = newList
}

func GetById(id int) Client {
	return filter(id)
}

func filter(id int) Client {

	var returnClient Client

	for x := 0; x < len(listClient); x++ {
		if listClient[x].Id == id {
			returnClient = listClient[x]
			return returnClient
		}
	}

	return returnClient
}

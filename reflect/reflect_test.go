package reflect

import "testing"

func TestCreateQuery(t *testing.T) {
	type person struct {
		Id     uint64 `json:"id"`
		Name   string `json:"name"`
		Gender int    `json:"gender"`
	}
	type animal struct {
		Id uint64 `json:"id"`
		T  string `json:"t"`
	}

	t.Log(CreateQuery(person{Id: 1, Name: "Scotty", Gender: 1}))
	t.Log(CreateQuery(animal{Id: 1, T: "Cat"}))

}

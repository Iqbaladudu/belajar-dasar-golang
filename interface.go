package main

import "fmt"

/**
=> Merupakan tipe data abstrak, tidak memiliki interface langsung
=> Sebuah interface berisikan defenisi-defenisi method
=> Biasanya interface digunakan untuk kontrak
=> Setiap tipe data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut
*/

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func SayHai(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var ima Person
	ima.Name = "Irma Gustia"
	SayHai(ima)

	cat := Animal{
		Name: "Meong",
	}
	SayHai(cat)
}

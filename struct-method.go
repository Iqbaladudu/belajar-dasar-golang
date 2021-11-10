package main

import "fmt"

type Konsumen struct {
	Name, Address string
	age           int
}

func (konsum Konsumen) SayHi(name string) {
	fmt.Println("Hello", name, "My Name is", konsum.Name)
}

func main() {
	ima := Konsumen{
		Name: "Irma Gustia",
	}
	ima.SayHi("Iqbal")
}

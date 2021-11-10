package main

import "fmt"

func main() {
	var name string
	name = "Muhammad Iqbal"
	fmt.Println(name)

	// Mengubah isi variabel
	name = "Irma Gustia"
	fmt.Println(name)

	// or
	var sister = "Irene"
	fmt.Println(sister)

	var age int8 = 20
	fmt.Println(age)

	// Without var
	country := "Indonesia"
	fmt.Println(country)

	// Multiple Variable
	var (
		firstName string
		lastName string
	)
	firstName = "Muhammad"
	lastName = "Iqbal"
	fmt.Println(firstName, lastName)
}
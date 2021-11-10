package main

import "fmt"

func getFullName() (string, string, int8) {
	return "Muhammad", "Iqbal", 10
}

func main() {
	firstName, lastName, age := getFullName()
	fmt.Println(firstName, lastName, age)

	// Ignoring return value
	_, _, umur := getFullName()
	fmt.Println(umur)
}

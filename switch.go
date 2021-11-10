package main

import "fmt"

func main() {
	name := "Iqbal"
	switch name {
	case "Iqbal":
		fmt.Println("Hello Iqbal")
	case "Irma":
		fmt.Println("Hello Mine")
	default:
		fmt.Println("Hello Hamba Allah")
	}

	// Short Statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Cocok")
	}
}

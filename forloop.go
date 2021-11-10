package main

import "fmt"

func main() {
	// counter := 1
	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke: ", counter)
	// 	counter++
	// }

	// For dengan statement
	// Init statement => dieksekusi sebelum for loop
	// Post statement => dieksekusi di akhir tiap perulangan
	// for init; expression; post {}

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	// Cara manual
	slice := []string{"Muhammad", "Iqbal", "Adudu"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// Menggunakan for range
	for index, value := range slice {
		fmt.Println(index, value)
	}

	// Tanpa index
	// _ => memberitahu golang kalau variabelnya tidak akan digunakan
	for _, value := range slice {
		fmt.Println(value)
	}

	// Looping in map
	person := make(map[string]string)
	person["name"] = "Iqbal"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, value)
	}
}

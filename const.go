package main

import "fmt"

func main() {
	// Konstan harus langsung diisi dengan value
	// Tidak wajib digunakan di Go, tidak akan error
	const firstName string = "Muhammad"
	const lastName = "Iqbal"
	const value = 1000
	fmt.Println(firstName, lastName, value)

	// Multiple const
	const (
		nama string = "Muhammad Iqbal"
		alamat string = "Indonesia"
		umur int8 =21
	)

}
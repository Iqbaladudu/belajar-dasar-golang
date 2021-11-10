package main

import "fmt"

/**
Struct => template data yang digunakan untuk menggabungkan NOL atau LEBIH tipe data yang berbeda dalam satu kesatuan. Struct biasanya representasi data dalam program aplikasi yang kita buat, data di dalam struct disimpan di dalam field.
Sederhananya, struct adalah kumpulan dari field
=> Struct adalah template data/prototype
=> Tidak bisa langsung digunakan
=> Kita bisa membuat data/object dari struct yang sudah kita buat
*/

type Customer struct {
	// Best practice-nya diawali huruf kapital
	Name, Address string
	Age           int
}

func main() {
	var iqbal Customer
	iqbal.Name = "Muhammad Iqbal"
	iqbal.Address = "Mesir"
	iqbal.Age = 21
	fmt.Println(iqbal)
	fmt.Println(iqbal.Name, iqbal.Address, iqbal.Age)

	// Struct Literals
	ima := Customer{
		Name:    "Irma Gustia",
		Address: "Turki",
		Age:     21,
	}
	fmt.Println(ima)
}

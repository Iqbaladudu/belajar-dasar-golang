package main

import "fmt"

func main() {
	// Tipe datanya harus sama
	// Daya Tampungnya tidak bisa ditambah setelah Array dibuat
	// Kita perlu menentukan jumlah data yang akan ditampung Array
	// var names [jumlahData]<tipeData>
	var names [3]string
	names[0] = "Muhammad"
	names[1] = "Iqbal"
	names[2] = "Irma"
	fmt.Println(names)
	fmt.Println(names[0], names[1], names[2])

	// Membuat Array saat deklarasi variabel
	var values = [3]int8{
		10,
		76,
		90,
	}
	fmt.Println(values)
	fmt.Println(len(values))
}
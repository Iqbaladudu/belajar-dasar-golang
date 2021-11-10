package main

import "fmt"

// jika semua tipe data returnya sama, tambahkan di tipe datanya di akhir variable saja
func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Muhammad"
	middleName = "Iqbal"
	lastName = "Adudu"

	// Tidak wajib disebutkan variabel returnya
	return firstName, middleName, lastName
}

func main() {
	namaDepan, namaTengah, _ := getCompleteName()
	fmt.Println(namaDepan, namaTengah)
}

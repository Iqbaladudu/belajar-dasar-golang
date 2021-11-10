package main

import "fmt"

func main() {
	var n32 int32 = 100000
	var n64 int64 = int64(n32)
	var n8 int8 = int8(n64)

	fmt.Println(n32)
	fmt.Println(n64)
	fmt.Println(n8)

	// string

	var name = "Iqbal"
	var e byte = name[0]
	var eString string = string(e)
	fmt.Println(name)
	fmt.Println(e)
	// Konversi byte ke string
	fmt.Println(eString)
}
package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool
	var noKTPIqbal NoKTP = "1233"
	fmt.Println(noKTPIqbal)
	var marriedStatus Married = false
	fmt.Println(marriedStatus)
}
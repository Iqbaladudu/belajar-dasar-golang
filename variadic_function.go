package main

import "fmt"

// args dalam Python
// cuma bisa satu
// hanya bisa di variabel paling belakang
func sumAll(numbers ...int) (total int) {
	total = 0
	for _, value := range numbers {
		total += value
	}

	return
}

func main() {
	total := sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(total)

	// Slice parameter
	numbers := []int{10, 65, 36, 57, 47, 46, 27}
	total = sumAll(numbers...)
	fmt.Println(total)
}

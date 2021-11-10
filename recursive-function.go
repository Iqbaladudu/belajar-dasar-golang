package main

import "fmt"

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursice(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursice(value-1)
	}
}

func main() {
	loop := factorialLoop(5)
	fmt.Println(loop)

	recursive := factorialRecursice(5)
	fmt.Println(recursive)
}

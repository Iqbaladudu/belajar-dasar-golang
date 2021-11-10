package main

import "fmt"

func main() {
	a := 10
	b := 70
	result1 := a > b
	c := 70
	d := 10
	result2 := c > d
	fmt.Println(result1 && result2)
	fmt.Println(result1 || result2)
}
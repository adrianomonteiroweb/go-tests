package main

import "fmt"

func main() {
	fmt.Println(Sum(1, 2, 3))
	fmt.Println(multiply(10, 10))
}

func Sum(i ...int) int {
	total := 0

	for _, v := range i {
		total += v
	}

	return total
}

func multiply(i ...int) int {
	total := 1

	for _, v := range i {
		total *= v
	}

	return total
}
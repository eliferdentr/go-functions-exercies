package main

import "fmt"

func main() {
	sum := sumUp(1, 2, 3, 4, 5)
	fmt.Println(sum)
	numbers := [] int {1, 10, 15, 40, -5}

	fmt.Println(sumUp(1, numbers...))
}

func sumUp(startingValue int, numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
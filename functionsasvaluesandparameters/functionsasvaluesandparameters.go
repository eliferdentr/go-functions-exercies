package functionsasvaluesandparameters

import "fmt"

type transformFn func (int) int

func main() {

	numbers := []int{1, 2, 3, 4}
	// transformNumbers(&numbers, double)
	// fmt.Println(numbers)
	// transformNumbers(&numbers, triple)
	// fmt.Println(numbers)
	moreNumbers := [] int {5, 6, 7 ,8}
	
	transformFn1 := getTransformerFunction(&numbers)
	transformFn2 := getTransformerFunction(&moreNumbers)
	

	transformNumbers(&numbers, transformFn1)
	transformNumbers(&moreNumbers, transformFn2)

	fmt.Println(numbers)
	fmt.Println(moreNumbers)
}

func transformNumbers(numbers *[]int, transform transformFn) {
	for index, value := range *numbers {
		(*numbers)[index] = transform(value)
	}
}

func getTransformerFunction (numbers *[] int) transformFn{
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
	
}

func double(number int) int {
	return number * 2
}

func triple (number int) int {
	return number * 3
}
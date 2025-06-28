package functionsarevalues

import "fmt"

type numberTransformer func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	moreNumbers := []int{6, 7, 8, 9, 10}
	// Demonstrating the use of a function as a parameter
	// and a function that returns a function
	// to transform numbers in a slice.
	doubledNumbers := transformNumbers(&numbers, double)
	tripledNumbers := transformNumbers(&numbers, triple)

	fmt.Println("Original numbers:", numbers)
	fmt.Println("Doubled numbers:", doubledNumbers)
	fmt.Println("Tripled numbers:", tripledNumbers)

	// Demonstrating the use of a function that returns a function
	transformerFunction1 := getNumberTransformerFunction(&numbers)
	transformerFunction2 := getNumberTransformerFunction(&moreNumbers)

	// Using the returned function to transform numbers
	transformedNumbers := transformNumbers(&numbers, transformerFunction1)
	moreNumbersTransformed := transformNumbers(&moreNumbers, transformerFunction2)

	fmt.Println("Transformed numbers using function from numbers:", transformedNumbers)
	fmt.Println("Transformed moreNumbers using function from moreNumbers:", moreNumbersTransformed)

}

func transformNumbers(numbers *[]int, transform numberTransformer) []int {
	doubled := []int{}

	for _, val := range *numbers {
		doubled = append(doubled, transform(val))
	}
	return doubled
}

func getNumberTransformerFunction(numbers *[]int) numberTransformer {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

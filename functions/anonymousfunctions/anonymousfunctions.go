package anonymousfunctions

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	double := createTransformerFunction(2)
	triple := createTransformerFunction(3)

	// Demonstrating anonymous function as a parameter
	// and a function that returns a function
	// to transform numbers in a slice.
	transformed := transformNumbers(&numbers, func(n int) int {
		return n * 2
	})

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(transformed)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformerFunction(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

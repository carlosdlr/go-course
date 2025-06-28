package main

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := sumup(1, 10, 15)
	println("Sum of numbers:", sum)

	anotherSum := sumup(1, numbers...)
	println("Sum of numbers with starting value:", anotherSum)

}

// variadic function that takes a starting value and a variable number of integers
// and returns the sum of all the integers including the starting value.
// This function demonstrates the use of variadic parameters in Go.
// Variadic functions allow you to pass a variable number of arguments to a function.
// The first parameter is the starting value, and the rest are the numbers to be summed.
// The function returns the sum of all the numbers.
func sumup(startingValue int, numbers ...int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

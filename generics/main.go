package main

func main() {
	result := add(5, 10)
	println(result) // Output: 15

	resultFloat := add(5.5, 10.5)
	println(resultFloat) // Output: 16

	resultString := add("Hello, ", "World!")
	println(resultString) // Output: Hello, World!
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}

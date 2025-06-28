package recursion

func main() {
	fact := factorial(5)
	println(fact)

}

// factorial of 5 => 5 * 4 * 3 * 2 * 1 = 120
// factorial of 0 => 1
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
	// result := 1
	// for i := 1; i <= n; i++ {
	// 	result *= i
	// }
	// return result
}

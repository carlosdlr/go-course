package lists

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println("Prices:", prices[0:1])
	prices[1] = 9.99 // updating the second element of the slice

	prices = append(prices, 5.99, 12.99, 29.99, 100.10) // append returns a new slice with the new values
	prices = prices[1:]
	fmt.Println("Updated Prices:", prices)

	discountPrices := []float64{5.99, 7.99, 3.99, 2.99}
	prices = append(prices, discountPrices...) // unpacking the slice
	fmt.Println("All Prices:", prices)
}

// 	var names [4]string = [4]string{"A book"}
// 	prices := [4]float64{10.99, 20.49, 5.99, 15.00}
// 	fmt.Println("Prices:", prices)

// 	names[2] = "A pen"

// 	fmt.Println(names)

// 	fmt.Println(prices[2])

// 	featuredPrices := prices[1:] // slice of prices from index 1 to 2. slices are a reference from a range of an array
// 	featuredPrices[0] = 18.99    // updating the first element of the slice
// 	highestPrices := featuredPrices[:1]
// 	fmt.Println(len(featuredPrices), cap(featuredPrices))
// 	fmt.Println(len(highestPrices), cap(highestPrices))
// 	fmt.Println(prices)

// 	highestPrices = highestPrices[:3]
// 	fmt.Println(highestPrices)
// 	fmt.Println(len(highestPrices), cap(highestPrices))

// }

package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println("Map contents:", m)
}

func main() {
	userNames := make([]string, 2, 5) // create a slice with a length of 2 and a capacity of 5

	userNames[0] = "Carlos"

	userNames = append(userNames, "Jane")
	userNames = append(userNames, "Alice", "Bob")
	userNames = append(userNames, "Charlie", "Dave") // appending more elements than the initial capacity

	fmt.Println("Usernames:", userNames)

	courseRatings := make(floatMap, 3)
	courseRatings["Go Programming"] = 4.5
	courseRatings["Python Programming"] = 4.7
	courseRatings["Java Programming"] = 4.2
	// adding a new course rating that is beyond the initial capacity cause memory reallocation
	courseRatings["JavaScript Programming"] = 4.8
	courseRatings.output()

	//fmt.Println("Course Ratings:", courseRatings)
	for index, value := range userNames {
		fmt.Printf("User %d: %s\n", index, value)
	}

	for key, value := range courseRatings {
		fmt.Printf("User %s: %f\n", key, value)
	}
}

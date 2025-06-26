package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	//1)
	hobbies := [3]string{"Reading", "Cycling", "Cooking"}
	fmt.Println("Hobbies:", hobbies)

	//2)
	fmt.Println("First Hobby:", hobbies[0])
	fmt.Println(hobbies[1:3])

	//3)
	mainHobbies := hobbies[0:2] // First way to create a slice
	fmt.Println("Main Hobbies (first way):", mainHobbies)

	//4)
	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3] // Second way to create a slice
	fmt.Println("Main Hobbies (second way):", mainHobbies)

	//5)
	courseGoals := []string{"Learn Go", "Build a web app"}
	fmt.Println("Course Goals:", courseGoals)

	//6)
	courseGoals[1] = "Master Go"
	courseGoals = append(courseGoals, "Contribute to open source")
	fmt.Println("Updated Course Goals:", courseGoals)

	//7)
	products := []Product{
		{"1", "Go Book", 29.99},
		{"2", "Go Course", 199.99},
	}

	fmt.Println("Products:", products)

	newProduct := Product{"3", "Go Conference Ticket", 99.99}
	products = append(products, newProduct)
	fmt.Println("Updated Products:", products)

}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.

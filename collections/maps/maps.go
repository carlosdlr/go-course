package maps

import "fmt"

func main() {
	websites := map[string]string{
		"google":        "https://www.google.com",
		"github":        "https://www.github.com",
		"stackoverflow": "https://www.stackoverflow.com",
		"golang":        "https://www.golang.org",
	}

	fmt.Println("Websites:", websites)
	fmt.Println("Google URL:", websites["google"])

	// mutating the map
	websites["linkedin"] = "https://linkedin.com"
	fmt.Println("Updated Websites:", websites)

	delete(websites, "google") // deleting an element from the map
	fmt.Println("After Deletion:", websites)

}

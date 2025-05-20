package main

import "fmt"

func main() {
	//Creating and initializing a map
	ages := map[string]int{
		"Jason":  47,
		"Angela": 42,
		"Sadie":  14,
	}

	fmt.Println("Initial map:", ages)

	//Acessing values
	fmt.Println("Jason's age:", ages["Jason"])

	//Updating a value
	ages["Angela"] = 43
	fmt.Println("Updated Angela's age:", ages["Angela"])

	//Adding a new key-value pair
	ages["Sophia"] = 22
	fmt.Println("Added Sophia:", ages)

	//Checking if a key exists
	val, exists := ages["Bob"]
	if exists {
		fmt.Println("Bob's age is:", val)
	} else {
		fmt.Println("Bob is not on the map.")
	}

	//Deleting a key
	delete(ages, "Alice")
	fmt.Println("Removed Alice:", ages)

	//Looping through a map
	fmt.Println("Listing all names and ages:")
	for name, age := range ages {
		fmt.Printf("- %s: %d years old\n", name, age)
	}
}

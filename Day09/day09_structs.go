package main

import "fmt"

//define a struct type
type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	//Create an instance of Person
	jason := Person{
		Name: "jason",
		Age:  47,
		City: "Salt Lake City",
	}

	// Accesising fields
	fmt.Println("Name:", jason.Name)
	fmt.Println("Age:", jason.Age)
	fmt.Println("City:", jason.City)

	// Modifying a field
	jason.City = "Cedar City"
	fmt.Println("Updated City:", jason.City)

	// Anouther instance
	alice := Person{"Alice", 29, "Denver"}
	fmt.Printf("Anouther Person: %+v\n", alice)

	// Passing struct to a function
	printGreeting(jason)
	printGreeting(alice)
}

// A function that takes a Person struct
func printGreeting(p Person) {
	fmt.Printf("Hello, %s from %s! You're %d years old.\n", p.Name, p.City, p.Age)
}

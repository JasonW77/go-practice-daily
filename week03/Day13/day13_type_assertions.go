package main

import "fmt"

// A function that accepts an empty interface (any type)
func describe(i interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", i, i)

	// Type assertion: check if 'i' is a string
	str, ok := i.(string)
	if ok {
		fmt.Println("It's a string:", str)
	} else {
		fmt.Println("Not a string.")
	}

	//Type switch: determine and respond to the actual type
	switch v := i.(type) {
	case int:
		fmt.Println("It's an integer:", v)
	case string:
		fmt.Println("It's a string (via switch):", v)
	default:
		fmt.Println("Unkown type.")
	}
}

func main() {
	describe("hello")
	describe(42)
	describe(3.14)
}

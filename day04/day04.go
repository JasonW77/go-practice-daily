package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Day 4: Functions and Error Handling")

	//call add function
	result := add(5, 7)
	fmt.Printf("5 + 7 = %d\n", result)

	//call divide function (with valid input)
	quotient, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", quotient)
	}

	//call divide function (with valid input)
	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// add returns the sum of the two integers
func add(a int, b int) int {
	return a + b
}

// divide returns the result of a division and an error if divisor is 0
func divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

package main

import "fmt"

func cleanup() {
	fmt.Println("Cleanup function executed with defer.")
}

func mayPanic(divisor int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Starting risky operation...")
	if divisor == 0 {
		panic("division by zero")
	}
	fmt.Println("Result:", 10/divisor)
}

func main() {
	defer cleanup()

	mayPanic(2) // Safe call
	mayPanic(0) // Will panic but be recovered

	fmt.Println("Program continues after recovery.")
}

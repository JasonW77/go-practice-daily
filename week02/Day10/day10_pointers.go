package main

import "fmt"

func main() {
	x := 10
	y := &x // y is a pointer to x

	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", y)
	fmt.Println("Value at address y:", *y)

	//modify x using the pointer
	*y = 20
	fmt.Println("Updated value of x:", x)

	//demonstrate pointer behavior in functions
	num := 5
	fmt.Println("Before increment:", num)
	increment(&num)
	fmt.Println("After increment:", num)

	// Value vs Refrence
	a := 100
	b := a
	b = 200
	fmt.Println("a:", a, "| b:", b) // a remains unchanged

	c := 100
	d := &c
	*d = 200
	fmt.Println("c:", c, "| *d:", *d) // c is changed
}

func increment(n *int) {
	*n = *n + 1
}

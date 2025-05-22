package main

import "fmt"

func main() {
	fmt.Println("Day 3: Loops, Arrays, and Slices")

	//Array: fixed length
	var numbers [5]int = [5]int{10, 20, 30, 40, 50}
	fmt.Println("Array contents:")
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Index %d: %d\n", i, numbers[i])
	}

	//Slice: dynamic length
	names := []string{"Jason", "Mira", "Toby", "Elena"}
	fmt.Println("\nSlice contents:")
	for index, name := range names {
		fmt.Printf("Name %d: %s\n", index, name)
	}

	//Adding to the slice
	names = append(names, "Nova")
	fmt.Println("\nUpdated slice:")
	for _, name := range names {
		fmt.Println(name)
	}
}

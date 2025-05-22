package main

import "fmt"

func main() {
	// Arrays: fixed length
	var nums [5]int
	nums[0] = 10
	nums[1] = 20
	nums[2] = 30

	fmt.Println("Array nums:", nums)
	fmt.Println("Length of nums:", len(nums))

	// Slices: dynamic size
	scores := []int{85, 90, 78}
	fmt.Println("Slice scores:", scores)

	// Apend to slice
	scores = append(scores, 95)
	fmt.Println("After append:", scores)

	// Slice anouther slice
	subScores := scores[1:3]
	fmt.Println("Sub-slice (index 1 to 2):", subScores)

	// Length and capacity
	fmt.Printf("Length: %d, Capacity: %d\n", len(scores), cap(scores))

	// Copying slices
	newScores := make([]int, len(scores))
	copy(newScores, scores)
	fmt.Println("Copied scores:", newScores)
}

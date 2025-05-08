package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Println("Welcome to day 2!")
	fmt.Print("What is your name? ")
	fmt.Scanln(&name)

	fmt.Print("How old are you? ")
	fmt.Scanln(&age)

	fmt.Printf("Nice to meet you, %s! You are %d years old. \n", name, age)

	if age < 18 {
		fmt.Println("You're still a young learner!")
	} else if age >= 18 && age < 60 {
		fmt.Println("Perfect age to learn Go.")
	} else {
		fmt.Println("You're a wise elder mastering Go!")
	}
}

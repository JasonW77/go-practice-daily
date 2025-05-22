package main

import "fmt"

//Define a struct named person
type Person struct {
	name string
	age  int
}

//Method: greet prints a greeting from the person
func (p Person) greet() {
	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", p.name, p.age)
}

//Method with pointer reciever: birthday increments the person's age
func (p *Person) birthday() {
	p.age++
	fmt.Printf("Happy Birthday, %s! You are now %d.\n", p.name, p.age)
}

func main() {
	fmt.Println("Day 5: Structs and Methods")

	//Create a new person
	jason := Person{name: "Jason", age: 47}

	//call methods
	jason.greet()
	jason.birthday()
	jason.greet()
}

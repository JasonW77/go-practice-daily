package main

import "fmt"

//define an interface
type Speaker interface {
	Speak()
}

//define types that implement the interface
type Person struct {
	Name string
}

func (p Person) Speak() {
	fmt.Printf("Hi, I'm %s and I like Go!\n", p.Name)
}

type Robot struct {
	Model string
}

func (r Robot) Speak() {
	fmt.Printf("Beep boop. Model %s online.\n", r.Model)
}

//Function that accepts any  type that implements Speaker
func announce(s Speaker) {
	s.Speak()
}

func main() {
	p := Person{Name: "Jason"}
	r := Robot{Model: "XJ9"}

	//interface in action
	announce(p)
	announce(r)
}

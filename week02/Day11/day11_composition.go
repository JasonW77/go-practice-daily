package main

import "fmt"

//base struct
type Animal struct {
	Name string
	Age  int
}

func (a Animal) Speak() {
	fmt.Printf("%s makes a sound. \n", a.Name)
}

//Embedded struct
type Dog struct {
	Animal
	Breed string
}

func (d Dog) Speak() {
	fmt.Printf("%s barks. (Breed: %s)\n", d.Name, d.Breed)
}

func main() {
	d := Dog{
		Animal: Animal{Name: "Rex", Age: 4},
		Breed:  "German Sheperd",
	}

	// Access promoted fields and methods
	d.Speak()        //Dogs method overides Animal's
	d.Animal.Speak() //Call base method explicitly
	fmt.Println("Age:", d.Age)
}

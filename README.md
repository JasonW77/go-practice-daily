# Go Practice Daily #

## Progress Overview ##

This repository tracks my daily Go programming practice. Each day, I will write and commit code to reinforce learning concepts, from basic syntax to more advanced Go features.


--- 

### Goals ###
- Gain familiarity with Go syntax and idioms.
- Build a solid foundation for developing Go applications.
- Track progress day-by-day, with a focus on consistent learning.

---

## Repository Structure ##
- Each day's practice is organized in a seperate folder (`day01`, `day02`, etc.) for clarity and progress tracking.

---

## Progress ##

**Day01** 
- ***Basic Go syntax, `main()` function, variables, and printing to the terminal***
- Set up Go development environment in VSCode.
- Created first Go program using fmt.Println() and fmt.Printf().

- Practiced:
    - Declaring variables (var age int = 47)
    - Short variable declarations (language := "Golang")
    - String interpolation using fmt.Printf()
    - Learned basic syntax structure of a Go program:
        - package main
        - func main()
    - Reinforced using the Go CLI: go run main.go

**Day02** 
- ***User input, conditionals, and basic types***
- Basic arithmetic operations: +, -, *, /, %
- Control flow using if, else if, else
- Logical operators: &&, ||, !
- Explored user input with fmt.Scan()
- Built a small CLI program to evaluate conditions and print responses.
- Strengthened understanding of Go's type system and variable handling.
- Got more comfortable with Go formatting and indentation.

**Day03** 
- ***Loops, arrays, slices, and iteration***
- Learned the difference between arrays and slices.
- Practiced looping with `for` and `range`.
- Used `append()` to grow a slice dynamically.

**Day04** 
- ***Functions and error handling***
- Wrote basic functions in Go using `func`.
- Practiced returning multiple values (result + error).
- Learned how to use `errors.New()` to return custom errors.

**Day05** 
- ***Structs and methods*** 
- Defined custom types with `struct`
- Attached methods to structs using value and pointer receivers
- Practiced modifying struct fields via methods

**Day 06** 
- ***Maps and basic I/O***
- Learned to use `map[string]int` to track word frequency.
- Practiced reading and writing files using `os` and `bufio`.
- Used `strings` to clean and process text.

**Day 07**
- ***Arrays and Slices***
- Declared arrays and initialized elements
- Learned the difference between arrays and slices
- Used `append()`, `copy()`, slicing syntax, and `make()`
- Checked slice `len()` and `cap()`
- Extracted sub-slices using slicing syntax

## Day 08 - Maps in Go

### Topics Covered:
- Declaring and initializing maps
- Accessing and modifying values
- Adding and deleting keys
- Checking if a key exists
- Iterating over a map with `range`

## Day 09 – Structs in Go

### Topics Covered:
- Declaring and initializing structs
- Field access and modification
- Struct literals (named and positional)
- Passing structs to functions
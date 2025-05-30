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

**Day 01** 
***Basic Go syntax, `main()` function, variables, and printing to the terminal***
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

**Day 02** 
***User input, conditionals, and basic types***
- Basic arithmetic operations: +, -, *, /, %
- Control flow using if, else if, else
- Logical operators: &&, ||, !
- Explored user input with fmt.Scan()
- Built a small CLI program to evaluate conditions and print responses.
- Strengthened understanding of Go's type system and variable handling.
- Got more comfortable with Go formatting and indentation.

**Day 03** 
***Loops, arrays, slices, and iteration***
- Learned the difference between arrays and slices.
- Practiced looping with `for` and `range`.
- Used `append()` to grow a slice dynamically.

**Day 04** 
***Functions and error handling***
- Wrote basic functions in Go using `func`.
- Practiced returning multiple values (result + error).
- Learned how to use `errors.New()` to return custom errors.

**Day 05** 
***Structs and methods*** 
- Defined custom types with `struct`
- Attached methods to structs using value and pointer receivers
- Practiced modifying struct fields via methods

**Day 06** 
***Maps and basic I/O***
- Learned to use `map[string]int` to track word frequency.
- Practiced reading and writing files using `os` and `bufio`.
- Used `strings` to clean and process text.

**Day 07**
***Arrays and Slices***
- Declared arrays and initialized elements
- Learned the difference between arrays and slices
- Used `append()`, `copy()`, slicing syntax, and `make()`
- Checked slice `len()` and `cap()`
- Extracted sub-slices using slicing syntax

**Day 08**
***Maps in Go***
- Declaring and initializing maps
- Accessing and modifying values
- Adding and deleting keys
- Checking if a key exists
- Iterating over a map with `range`

**Day 09**
***Structs in Go***
- Declaring and initializing structs
- Field access and modification
- Struct literals (named and positional)
- Passing structs to functions

**Day 10** 
***Pointers in Go***
- Getting the address of a variable with `&`
- Dereferencing with `*`
- Passing pointers to functions to modify original data
- The difference between value and reference types

**Day 11**
***Struct Composition & Embedded Types***
- Learned how Go uses composition instead of traditional inheritance.
- Practiced embedding one struct into another and promoting fields/methods.
- Overrode methods in embedded types and accessed parent methods explicitly.
- Code example: `day11_composition.go`

**Day 12**
***Interfaces in Go***
- Learned what interfaces are and how they define behavior.
- Practiced implementing interfaces with multiple struct types.
- Used interfaces as function parameters to create reusable and flexible code.
- Code example: `day12_interfaces.go`

**Day 13**
***Type Assertions & Type Switches***
- Practiced using the empty interface (`interface{}`) to accept any value.
- Used type assertions to safely extract concrete values from interfaces.
- Used type switches to handle different types dynamically.
- Code example: `day13_type_assertions.go`

**Day 14** 
***Error Handling in Go***
- Learned to return and check errors using Go’s standard `error` type.
- Used `errors.New` and `fmt.Errorf` for creating meaningful error messages.
- Wrote a custom `divide` function to demonstrate error handling with zero division.
- Code example: `week03/day14_errors.go`

**Day 15** 
***Defer, Panic, and Recover***
- Explored Go's mechanisms for cleanup and error recovery.
- Used `defer` to execute logic when a function returns.
- Triggered a `panic` and handled it with `recover`.
- Demonstrated how to write safe and robust code in the face of runtime errors.
- Code example: `week03/day15_defer_panic_recover.go`

Key Concepts:
- `defer` for deferred execution
- `panic` to stop the program unexpectedly
- `recover` to regain control from a panic

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

**Day 16**
***Making HTTP Requests***
- Learned how to perform basic HTTP GET requests using `net/http`.
- Handled response errors and printed status and body.
- Used `defer` to close the response body properly.
- Code example: `week03/day16_http_get.go`

Key Concepts:
- `http.Get`
- `io.ReadAll`
- `defer resp.Body.Close()`
- Basic error handling

**Day 17**
***Parsing JSON into Structs***
- Learned to use Go’s `encoding/json` package to decode JSON into native Go structs.
- Defined a `User` struct with JSON tags for correct field mapping.
- Parsed JSON strings using `json.Unmarshal`.
- Handled errors during decoding and printed structured output.
- Code example: `week04/Day17/day17_parse_json.go`

Key Concepts:
- `json.Unmarshal([]byte, &struct)`
- Struct tags: `json:"field_name"`
- Working with `string` and `[]byte` JSON inputs

**Day 18**
***Encoding Structs to JSON***
- Learned how to encode Go structs into JSON using `json.Marshal`.
- Used struct tags (`json:"field_name"`) to control JSON field names.
- Demonstrated how to pretty-print JSON using `json.MarshalIndent`.
- Wrote a small program that constructs a struct, encodes it to JSON, and prints the result.
- Code example: `week04/Day18/day18_encode_json.go`

Key Concepts:
- `json.Marshal(struct)`
- `json.MarshalIndent(struct, "", "  ")`
- Struct field visibility (must be exported to encode)

## upcomming projects:
Day 19 – Reading from JSON files
Day 20 – Writing to JSON files
Day 21 – Combining file I/O and JSON parsing
Day 22 – Build a mini JSON config reader
Day 23 – Review: build a small app using config + file persistence
Day 24 – Using os.Args for CLI arguments
Day 25 – Building a flag parser with flag package
Day 26 – Building a todo list CLI app (add task)
Day 27 – Expanding todo CLI (mark as done)
Day 28 – Save tasks to a file
Day 29 – Load tasks from file
Day 30 – Review + polish CLI app
Day 31 – Introduction to goroutines
Day 32 – Using channels for communication
Day 33 – Buffered vs unbuffered channels
Day 34 – The select statement
Day 35 – Goroutines with anonymous functions
Day 36 – Timeout with time.After
Day 37 – Review concurrency basics with a simulation
Day 38 – Worker pool pattern
Day 39 – Fan-in / Fan-out pattern
Day 40 – Rate limiting with time.Ticker
Day 41 – Mutexes and race conditions
Day 42 – Using sync.WaitGroup
Day 43 – Using sync.Once and sync.Map
Day 44 – Review: build a concurrent downloader
Day 45 – Writing basic unit tests with testing package
Day 46 – Table-driven tests
Day 47 – Benchmark tests
Day 48 – Code coverage
Day 49 – Mocks and interfaces
Day 50 – Using testify package
Day 51 – Review: test-driven small module
Day 52 – Writing a TCP server
Day 53 – Writing a TCP client
Day 54 – Echo server
Day 55 – Handle multiple clients
Day 56 – Using bufio with TCP
Day 57 – Graceful shutdown
Day 58 – Review project: TCP chat app
Day 59 – HTTP server basics
Day 60 – Handling routes
Day 61 – Query parameters and POST forms
Day 62 – Serving static files
Day 63 – Basic templating with html/template
Day 64 – Sessions and cookies
Day 65 – Build a simple blog engine
Day 66 – Gorilla Mux router
Day 67 – Building RESTful routes
Day 68 – JSON APIs with Gorilla
Day 69 – Middleware (logging, auth)
Day 70 – Using Gorilla sessions
Day 71 – Authentication with JWT
Day 72 – Build a full-featured REST API
Day 73 – Connecting to SQLite/Postgres
Day 74 – Basic SQL CRUD with database/sql
Day 75 – Using sqlx or gorm
Day 76 – Handling relationships
Day 77 – Database migrations
Day 78 – Error handling and transactions
Day 79 – Build: Todo app with DB backend
Day 80 – Understanding the Go memory model
Day 81 – Deep dive into slices and arrays
Day 82 – Exploring the reflect package
Day 83 – Writing a custom JSON marshaller
Day 84 – Go GC overview
Day 85 – Escape analysis and performance
Day 86 – Review: build a custom serializer
Day 87 – Factory pattern
Day 88 – Singleton pattern
Day 89 – Strategy pattern
Day 90 – Observer pattern
Day 91 – Command pattern
Day 92 – Adapter pattern
Day 93 – Review: use 3 patterns in a small project
Day 94 – Plan capstone project
Day 95 – Set up project structure
Day 96 – Implement core logic
Day 97 – Add API layer
Day 98 – Add persistence layer
Day 99 – Write tests and polish
Day 100 – Final review and push to GitHub portfolio
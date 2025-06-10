package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	user := User{
		Name:  "Jason Waters",
		Email: "jason@example.com",
		Age:   47,
	}

	// Compact JSON
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	fmt.Println("Compact JSON:")
	fmt.Println(string(jsonData))

	// Pretty JSON
	prettyJSON, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Println("Error pretty-printing JSON:", err)
		return
	}
	fmt.Println("\nPretty JSON:")
	fmt.Println(string(prettyJSON))
}

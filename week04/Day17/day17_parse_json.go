package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
}

func main() {
	jsonData := `{
		"name": "Jason Waters",
		"age": 47,
		"email": "jason@example.com",
		"is_active": true
	}`

	var user User

	err := json.Unmarshal([]byte(jsonData), &user)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println("Parsed JSON into Go struct:")
	fmt.Printf("Name: %s\nAge: %d\nEmail: %s\nActive: %t\n",
		user.Name, user.Age, user.Email, user.IsActive)
}

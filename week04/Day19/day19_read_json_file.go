package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	file, err := os.Open("week04/Day19/users.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("Users from JSON file:")
	for _, user := range users {
		fmt.Printf("Name: %s, Email: %s, Age: %d\n", user.Name, user.Email, user.Age)
	}
}

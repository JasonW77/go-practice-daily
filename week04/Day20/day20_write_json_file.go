package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	products := []Product{
		{ID: 1, Name: "Keyboard", Price: 29.99},
		{ID: 2, Name: "Mouse", Price: 14.95},
		{ID: 3, Name: "Monitor", Price: 199.99},
	}

	data, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	file, err := os.Create("week04/Day20/products.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Successfully wrote JSON to products.json")
}

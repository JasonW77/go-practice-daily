package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countWords(text string) map[string]int {
	wordCount := make(map[string]int)
	words := strings.Fields(text)
	for _, word := range words {
		clean := strings.ToLower(strings.Trim(word, ".,!?;:\""))
		wordCount[clean]++
	}
	return wordCount
}

func main() {
	fmt.Println("Day 6: Maps and Basic I/O")

	//Write to file
	file, err := os.Create("input.txt")
	if err != nil {
		fmt.Println("error creating file:", err)
		return
	}
	defer file.Close()

	sample := "Hello Go! Go is fast. Go is fun. Learning Go with Jason."
	_, err = file.WriteString(sample)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Read from file
	file, err = os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content string
	for scanner.Scan() {
		content += scanner.Text() + " "
	}

	// Word counting
	counts := countWords(content)
	fmt.Println("Word Counts:")
	for word, count := range counts {
		fmt.Printf("%s: %d\n", word, count)
	}
}

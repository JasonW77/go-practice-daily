package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://api.github.com"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("response Body:\n", string(body))
}

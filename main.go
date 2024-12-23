package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {

	http.HandleFunc("/calc", func(w http.ResponseWriter, r *http.Request) {

		text := r.URL.Query().Get("text")
		if text == "" {
			text = "default input"
		}

		if isAddOrSub(text) == "add" {
			numbers, _ := parseExpression(text)
			result := add(numbers[0], numbers[1])
			fmt.Fprintf(w, "Result: %d\n", result)
		} else if isAddOrSub(text) == "sub" {
			numbers, _ := parseExpression(text)
			result := sub(numbers[0], numbers[1])
			fmt.Fprintf(w, "Result: %d\n", result)
		}
	})
	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server closed")
	} else if err != nil {
		fmt.Printf("Error starting server %s\n", err)
		os.Exit(1)
	}

}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func parseExpression(text string) ([]int, []rune) {
	var numbers []int
	var chars []rune

	for _, char := range text {
		if char == '+' || char == '-' {
			chars = append(chars, char)
		} else {
			number, _ := strconv.Atoi(string(char))
			numbers = append(numbers, number)
		}
	}
	return numbers, chars
}

func isAddOrSub(text string) string {
	_, char := parseExpression(text)
	if char[0] == '+' {
		return "add"
	} else {
		return "sub"
	}
}

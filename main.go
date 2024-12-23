package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "got / request\n")
	io.WriteString(w, "This is a test\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server is running on port 8080")

	err := http.ListenAndServe(":8080", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server is shut down")
	} else if err != nil {
		fmt.Println("Server is not shut down")
		os.Exit(1)
	}
}

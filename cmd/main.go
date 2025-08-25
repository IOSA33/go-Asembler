package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go hello.asm")
		return
	}

	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %d", err)
		return
	}

}

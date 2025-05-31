package main

import (
	"fmt"
	"os"
)

func main() {
	paths := os.Args[1:]

	if len(paths) == 0 {
		fmt.Println("You need to provide at least one file path as an argument.")
		return
	}

	for _, path := range paths {
		content, err := os.ReadFile(path)

		if err != nil {
			fmt.Println("Error reading file:", err)
		}

		fmt.Println(string(content))
	}
}

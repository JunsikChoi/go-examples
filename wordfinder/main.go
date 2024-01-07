package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Error: Not enough arguments")
		fmt.Println("Usage: wordfinder [target word] ...[target files or patterns]")
		return
	}

	word := os.Args[1]
	files := os.Args[2:]

	fmt.Println("Target word: ", word)
	fmt.Println(files)
	GetFiles(files)
}

func GetFiles(files []string) {
	for _, path := range files {
		matches, err := filepath.Glob(path)
		if err != nil {
			fmt.Println("Error occurs during finding files", path)
			return
		}
		fmt.Printf("Filelist for %s\n", path)
		for _, match := range matches {
			fmt.Println(match)
		}
	}
}
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) == 4 && strings.HasPrefix(os.Args[1], "--output=") {

		filename := strings.TrimPrefix(os.Args[1], "--output=")

		input := os.Args[2]
		banner := os.Args[3] + ".txt"

		asciiArt := LoadBanner(input, banner)

		err := os.WriteFile(filename, []byte(asciiArt), 0644)
		if err != nil {
			fmt.Println(err)
		}

		return
	}

	if len(os.Args) < 2 {
		fmt.Println("error")
		return
	}

	banner := "standard.txt"

	if len(os.Args) == 3 {
		banner = os.Args[2] + ".txt"
	}

	input := os.Args[1]

	if input == "" {
		return
	}

	if input == "\\n" {
		fmt.Println()
		return
	}

	fmt.Print(LoadBanner(input, banner))
}

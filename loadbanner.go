package main

import (
	"os"
	"strings"
)

func LoadBanner(text string, banner string) string {
	data, err := os.ReadFile(banner)
	if err != nil {
		return "unsupported file"
	}

	lines := strings.Split(string(data), "\n")
	fontMap := make(map[rune][]string)

	for char := ' '; char <= '~'; char++ {
		begin := (int(char) - 32) * 9
		fontMap[char] = lines[begin+1 : begin+9]
	}

	var output strings.Builder

	word := strings.Split(text, "\\n")

	for _, words := range word {
		if words == "" {
			output.WriteString("\n")
			continue
		}

		for row := 0; row < 8; row++ {
			for _, char := range words {
				output.WriteString(fontMap[char][row])
			}
			output.WriteString("\n")
		}
	}

	return output.String()
}

package main

import (
	"fmt"
	"read_write/fileutils"
)

func main() {
	originalContent, err := fileutils.ReadTextFile("data/text.txt")

	if err != nil {
		panic(err)
	}

	fmt.Printf("Original Content\n%s\n", originalContent)

	multipliedContent := ""

	for i := 0; i < 5; i++ {
		multipliedContent += fmt.Sprintf("%s\n", originalContent)
	}

	err = fileutils.WriteTextFile("data/multiplied_text.txt", multipliedContent)

	if err != nil {
		panic(err)
	}

	multipliedContent, err = fileutils.ReadTextFile("data/multiplied_text.txt")

	fmt.Printf("\nMultiplied Content\n%s\n", multipliedContent)

}

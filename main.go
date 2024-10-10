package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the filename as a parameter.")
		return
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	var outputLines []string
	scanner := bufio.NewScanner(file)

	// Regular expressions for the required transformations
	checkedBoxPattern := regexp.MustCompile(`- \[x\]`)
	textPattern := regexp.MustCompile(`(.*) >: (.*)`)

	for scanner.Scan() {
		line := scanner.Text()

		// Replace "- [x]" with "- [ ]"
		line = checkedBoxPattern.ReplaceAllString(line, "- [ ]")

		// Replace "any text >: some text" with "any text >: ___"
		line = textPattern.ReplaceAllString(line, "$1 >: ___")

		outputLines = append(outputLines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Write the transformed content back to the file
	err = os.WriteFile(filename, []byte(strings.Join(outputLines, "\n")), 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}

	fmt.Println("File transformation completed successfully.")
}

package helpers

import (
	"bufio"
	"os"
	"strings"
)

func ReadFileWithSeparator(filename string, separator string) [][]string {
	fileContents := [][]string{}
	file, err := os.Open(filename)

	if err != nil {
		panic("File not found")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		fileContents = append(
			fileContents,
			strings.Split(scanner.Text(), separator),
		)
	}

	if err := scanner.Err(); err != nil {
		panic("Error with scanner")
	}

	return fileContents
}

func ReadFile(filename string) []string {
	fileContents := []string{}
	file, err := os.Open(filename)

	if err != nil {
		panic("File not found")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		fileContents = append(
			fileContents,
			scanner.Text(),
		)
	}

	if err := scanner.Err(); err != nil {
		panic("Error with scanner")
	}

	return fileContents
}

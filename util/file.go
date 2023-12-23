package util

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFileToLines(filePath string) (*bufio.Scanner, error) {
	fileInput, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(fileInput)
	fileScanner.Split(bufio.ScanLines)

	return fileScanner, nil
}

func PrintSplice(input []string) {
	for index, line := range input {
		fmt.Printf("[%d]: %s", index, line)
	}
}

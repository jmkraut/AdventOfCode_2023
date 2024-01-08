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

func ReadFileTo2DStringArray(filepath string) ([][]string, error) {
	fileInput, err := os.Open(filepath)

	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(fileInput)
	var stringArray [][]string

	for fileScanner.Scan() {
		line := fileScanner.Text()
		stringLine := []string{}

		for _, char := range line {
			stringLine = append(stringLine, string(char))
		}

		stringArray = append(stringArray, stringLine)
	}

	return stringArray, nil
}

func PrintStringSlice(input []string) {
	for index, line := range input {
		fmt.Printf("[%d]: %s", index, line)
	}
}

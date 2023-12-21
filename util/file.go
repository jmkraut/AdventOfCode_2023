package util

import (
	"fmt"
	"os"
	"strings"
)

func ReadFileToStringSplice(filename string) ([]string, error) {
	file, err := os.ReadFile(filename)
	check(err)

	return strings.Split(string(file), "\n"), nil
}

func PrintSpliceArray(input []string) {
	for index, line := range input {
		fmt.Printf("[%d]: %s", index, line)
	}
}

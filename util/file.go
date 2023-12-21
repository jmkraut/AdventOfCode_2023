package util

import (
	"bufio"
	"os"
)

func ReadFileToBuffer(filePath string) *bufio.Reader {
	file, err := os.Open(filePath)
	defer file.Close()

	check(err)

	return bufio.NewReader(file)
}

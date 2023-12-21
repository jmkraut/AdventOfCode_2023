package day_1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	DAY_1_FILEPATH = "resources\\day_1\\input.txt"
)

func SolveDay1() {
	fileInput, _ := os.ReadFile(DAY_1_FILEPATH)
	total := 0
	var stringbuilder strings.Builder

	for index, char := range fileInput {
		r := rune(char)

		if unicode.IsDigit(r) {
			stringbuilder.WriteRune(r)
		}

		if r == '\n' || index+1 == len(fileInput) {
			if stringbuilder.Len() >= 2 {
				stringbuilderFirstAndLast := string(stringbuilder.String()[0]) + string(stringbuilder.String()[stringbuilder.Len()-1])
				stringbuilderTotal, _ := strconv.Atoi(stringbuilderFirstAndLast)
				total += stringbuilderTotal
			} else {
				stringbuilder.WriteString(stringbuilder.String())
				stringbuilderTotal, _ := strconv.Atoi(stringbuilder.String())
				total += stringbuilderTotal
			}

			stringbuilder.Reset()
		}
	}

	fmt.Println(strconv.Itoa(total))
}

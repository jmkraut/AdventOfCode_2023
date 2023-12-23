package days

import (
	util "aoc2023/util"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

const (
	DAY_1_FILEPATH = "resources\\day_1\\input_test_part_2.txt"
)

func SolvePart1() {
	fileInput, _ := os.ReadFile(DAY_1_FILEPATH)
	total := 0
	var stringBuilder strings.Builder

	for index, char := range fileInput {
		charAsRune := rune(char)

		if unicode.IsDigit(charAsRune) {
			stringBuilder.WriteRune(charAsRune)
		}

		// If the index+1 is the same as the fileinput length we've hit the end of the file
		if charAsRune == '\n' || index+1 == len(fileInput) {
			if stringBuilder.Len() >= 2 {
				stringbuilderFirstAndLast := string(stringBuilder.String()[0]) + string(stringBuilder.String()[stringBuilder.Len()-1])
				stringbuilderTotal, _ := strconv.Atoi(stringbuilderFirstAndLast)
				total += stringbuilderTotal
			} else {
				// If it's less than 2 in length, we need to repeat the string
				stringBuilder.WriteString(stringBuilder.String())
				stringbuilderTotal, _ := strconv.Atoi(stringBuilder.String())
				total += stringbuilderTotal
			}

			stringBuilder.Reset()
		}
	}

	fmt.Println(strconv.Itoa(total))
}

func SolvePart1Alt() {
	fileScanner, _ := util.ReadFileToLines(DAY_1_FILEPATH)

	total := 0

	for fileScanner.Scan() {
		str := fileScanner.Text()
		str = regexp.MustCompile(`\D`).ReplaceAllString(str, "")
		subTotal := 0

		if len(str) > 1 {
			subTotal, _ = (strconv.Atoi(string(str[0]) + string(str[len(str)-1])))
		} else {
			subTotal, _ = (strconv.Atoi(string(str) + string(str)))
		}

		total += subTotal
	}

	fmt.Print(total)
}

func SolvePart2() {
	fileScanner, _ := util.ReadFileToLines(DAY_1_FILEPATH)

	total := 0
	// stringNumbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	// stringNumbers := []string{"one"}

	for fileScanner.Scan() {
		str := fileScanner.Text()

		var stringNumberResult []string
		// var leftMost, rightMost int

		// for _, number := range stringNumbers {
		stringNumberResult = regexp.MustCompile(`(e|t|h|o|u|i|w|r)`).FindAllString(str, -1)

		if len(stringNumberResult) > 0 {
			for _, result := range stringNumberResult {
				fmt.Println(result[0])
			}
		}
		// }

		// numberResult = regexp.MustCompile(`\d`).FindAllStringIndex(str, -1)
		// str := fileScanner.Text()
		// str = regexp.MustCompile(`\D`).ReplaceAllString(str, "")
		// subTotal := 0

		// if len(str) > 1 {
		// 	subTotal, _ = (strconv.Atoi(string(str[0]) + string(str[len(str)-1])))
		// } else {
		// 	subTotal, _ = (strconv.Atoi(string(str) + string(str)))
		// }

		// total += subTotal
	}

	fmt.Print(total)
}

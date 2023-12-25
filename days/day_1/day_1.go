package days

import (
	util "aoc2023/util"
	"fmt"
	"regexp"
	"strconv"
)

const (
	DAY_1_FILEPATH = "resources\\day_1\\input.txt"
)

func Solve() {
	fmt.Println(solvePart1())
	fmt.Println(solvePart2())
}

func solvePart1() string {
	fileScanner, _ := util.ReadFileToLines(DAY_1_FILEPATH)

	total := 0

	for fileScanner.Scan() {
		str := fileScanner.Text()
		str = regexp.MustCompile(`\D`).ReplaceAllString(str, "")

		total += getSubTotalFromString(str)
	}

	return fmt.Sprintf("Solution D1P1: %d", total)
}

func solvePart2() string {
	fileScanner, _ := util.ReadFileToLines(DAY_1_FILEPATH)

	total := 0

	for fileScanner.Scan() {
		str := fileScanner.Text()

		str = tokenizeString(str)
		str = regexp.MustCompile(`\D`).ReplaceAllString(str, "")
		total += getSubTotalFromString(str)
	}

	return fmt.Sprintf("Solution D1P2: %d", total)
}

func getSubTotalFromString(str string) int {
	var subTotal int
	if len(str) > 1 {
		subTotal, _ = (strconv.Atoi(string(str[0]) + string(str[len(str)-1])))
	} else {
		subTotal, _ = (strconv.Atoi(string(str) + string(str)))
	}

	return subTotal
}

func tokenizeString(str string) string {
	// Tokenize the string numbers to allow them to be countable but without
	// damaging strings that share a letter e.g. 'eightwothree'
	tokenizeMap := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

	for key, value := range tokenizeMap {
		str = regexp.MustCompile(key).ReplaceAllString(str, value)
	}

	return str
}

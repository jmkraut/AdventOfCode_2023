package day_2

import (
	"aoc2023/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	DAY_2_INPUT_FILEPATH       = "resources\\day_2\\input.txt"
	DAY_2_PART_1_TEST_FILEPATH = "resources\\day_2\\input_test_part_1.txt"
	DAY_2_PART_2_TEST_FILEPATH = "resources\\day_2\\input_test_part_2.txt"
)

func Solve() {
	fmt.Println(solvePart1())
	fmt.Println(solvePart2())
}

func solvePart1() string {
	fileScanner, _ := util.ReadFileToLines(DAY_2_INPUT_FILEPATH)

	gameIdTotal := 0
	gameIsImpossible := false

	for fileScanner.Scan() {

		str := fileScanner.Text()
		cubeRegex := regexp.MustCompile(`(\d{0,2}) (?:blue|green|red|blue;|green;|red;)`).FindAllStringSubmatch(str, -1)

		for i := 0; i < len(cubeRegex); i++ {
			cubeTotal, _ := strconv.Atoi(cubeRegex[i][1])

			impossibleBlue := strings.Contains(cubeRegex[i][0], "blue") && cubeTotal > 14
			impossibleGreen := strings.Contains(cubeRegex[i][0], "green") && cubeTotal > 13
			impossibleRed := strings.Contains(cubeRegex[i][0], "red") && cubeTotal > 12

			gameIsImpossible = impossibleBlue || impossibleGreen || impossibleRed

			if gameIsImpossible {
				break
			}
		}

		if !gameIsImpossible {
			gameId, _ := strconv.Atoi(regexp.MustCompile(`Game (\d{1,3}): `).FindStringSubmatch(str)[1])
			gameIdTotal += gameId
		}
	}

	return fmt.Sprintf("Solution D2P1: %d", gameIdTotal)
}

func solvePart2() string {
	fileScanner, _ := util.ReadFileToLines(DAY_2_PART_2_TEST_FILEPATH)

	powerOfCubesTotal := 0

	for fileScanner.Scan() {

		str := fileScanner.Text()
		cubeRegex := regexp.MustCompile(`(\d{0,2}) (?:blue|green|red|blue;|green;|red;)`).FindAllStringSubmatch(str, -1)

		largestBlue := 0
		largestGreen := 0
		largestRed := 0

		for i := 0; i < len(cubeRegex); i++ {
			cubeTotal, _ := strconv.Atoi(cubeRegex[i][1])

			currentBlueIsGreaterThanPrevious := strings.Contains(cubeRegex[i][0], "blue") && cubeTotal > largestBlue
			currentGreenIsGreaterThanPrevious := strings.Contains(cubeRegex[i][0], "green") && cubeTotal > largestGreen
			currentRedIsGreaterThanPrevious := strings.Contains(cubeRegex[i][0], "red") && cubeTotal > largestRed

			if currentBlueIsGreaterThanPrevious {
				largestBlue = cubeTotal
			}
			if currentGreenIsGreaterThanPrevious {
				largestGreen = cubeTotal
			}
			if currentRedIsGreaterThanPrevious {
				largestRed = cubeTotal
			}
		}

		powerOfCubesTotal += (largestBlue * largestRed * largestGreen)

	}

	return fmt.Sprintf("Solution D2P2: %d", powerOfCubesTotal)
}

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
}

func solvePart1() string {
	fileScanner, _ := util.ReadFileToLines(DAY_2_INPUT_FILEPATH)

	gameIdTotal := 0
	gameIsImpossible := false
	var cubeRegex [][]string

	for fileScanner.Scan() {

		str := fileScanner.Text()
		cubeRegex = regexp.MustCompile(`(\d{0,2}) (?:blue|green|red|blue;|green;|red;)`).FindAllStringSubmatch(str, -1)

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
			gameRegex := regexp.MustCompile(`Game (\d{1,3}): `).FindStringSubmatch(str)
			gameId, _ := strconv.Atoi(gameRegex[1])
			fmt.Printf("Game %d is possible!\n", gameId)
			gameIdTotal += gameId
		}
	}

	return fmt.Sprintf("Solution D2P1: %d", gameIdTotal)
}

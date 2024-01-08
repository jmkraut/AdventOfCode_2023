package day_3

import (
	"aoc2023/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	DAY_3_INPUT_FILEPATH       = "resources\\day_3\\input.txt"
	DAY_3_PART_1_TEST_FILEPATH = "resources\\day_3\\input_test_part_1.txt"
	DAY_3_PART_2_TEST_FILEPATH = "resources\\day_3\\input_test_part_2.txt"
)

func Solve() {
	fmt.Println(solvePart1())
	fmt.Println(solvePart2())
}

func solvePart1() string {
	schematicsArray, _ := util.ReadFileTo2DStringArray(DAY_3_INPUT_FILEPATH)

	var stringBuilder strings.Builder
	schematicsTotal := 0
	specialCharFoundAdjacent := false

	for row := 0; row < len(schematicsArray); row++ {
		for column := 0; column < len(schematicsArray[row]); column++ {
			schematicItem := schematicsArray[row][column]
			// check if the value is a number, if it is, we add it to the builder.
			// we also check if any value surrounding the number is adjacent to a special char
			if strings.ContainsAny(schematicItem, "0123456789") {
				stringBuilder.WriteString(schematicItem)

				if specialCharacterIsAdjacent(row, column, schematicsArray) {
					specialCharFoundAdjacent = true
				}
			}

			// if the value isn't a number, we've hit the end of our number and can grab the result
			if !strings.ContainsAny(schematicItem, "0123456789") {
				if specialCharFoundAdjacent {
					stringBuilderTotal, _ := strconv.Atoi(stringBuilder.String())
					schematicsTotal += stringBuilderTotal
					stringBuilder.Reset()
					specialCharFoundAdjacent = false
				} else {
					stringBuilder.Reset()
				}
			}
		}
	}

	return fmt.Sprintf("Solution D3P1: %d", schematicsTotal)
}

func specialCharacterIsAdjacent(row int, column int, schematicArray [][]string) bool {
	// get our boundaries to check, ensuring we don't go outside the arrays
	rowStartIndex := int(math.Max(0.0, float64(row-1)))
	rowFinishIndex := int(math.Min(float64(row+2), float64(len(schematicArray))))
	columnStartIndex := int(math.Max(0.0, float64(column-1)))
	columnFinishIndex := int(math.Min(float64(column+2), float64(len(schematicArray[row]))))

	for i := rowStartIndex; i < rowFinishIndex; i++ {
		for j := columnStartIndex; j < columnFinishIndex; j++ {
			schematicItem := schematicArray[i][j]
			if strings.ContainsAny(schematicItem, "!@#$%^&*()_+-=/") {
				return true
			}
		}
	}

	return false
}

func solvePart2() string {
	return "complete me"
}

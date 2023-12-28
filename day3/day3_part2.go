package main

import (
	"unicode"
)

type Coordinate struct {
	row, col int
}

func getSumOfGearRatios(s string) int {
	gearRatiosSum := 0
	mapOfPartNumbers := getMapOfPartNumbers(s)

	matrix := convertStringToMatrix(s)
	n := len(matrix)

	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			s := string(matrix[row][col])
			if s == "*" {
				gearRatiosSum += getGearRatioOfGear(matrix, mapOfPartNumbers, row, col)
			}
		}
	}
	return gearRatiosSum
}

/*
getGearRatioOfGear checks the adjecent cells in a clock wise mode starting from the cell at left.
Check getGearRatioOfGear.png.
It will check in the following order (where -1 or +1 is relative to where the gear is):
[0][-1] then [-1][-1] then [-1][0] then [-1][+1] ... and lastly [+1][-1].
A gear is any * symbol that is adjacent to exactly two part numbers.

Meaning of A,B,C (positions we need to check relative to the gear).
Given as input example:
7 . .
. * .
3 5 .

it means:
7   A  B
.   *  .
D   C  .

which using the relative indexes, means:
[-1][-1] [-1][0] [-1][+1]

[0][-1]     *    [0][+1]

[+1][-1] [+1][0] [+1][+1]
*/
func getGearRatioOfGear(matrix [][]rune, mapOfPartNumbers map[Coordinate]int, row, col int) int {
	n := len(matrix)
	nAdjecentPartNumbers := 0
	gearRatio := 1
	var x, y int

	// check left cell
	x = row
	y = col - 1
	if y >= 0 {
		c := Coordinate{x, y}
		fromMap, present := mapOfPartNumbers[c]
		if present {
			gearRatio *= fromMap
			nAdjecentPartNumbers++
		}
	}

	//check above row
	x = row - 1
	y = col - 1
	shouldCheckA := true
	if x >= 0 && y >= 0 {
		c := Coordinate{x, y}
		fromMap, present := mapOfPartNumbers[c]
		if present {
			shouldCheckA = false
			gearRatio *= fromMap
			nAdjecentPartNumbers++
		}
	}

	x = row - 1
	y = col
	shouldCheckB := true
	if x >= 0 && unicode.IsDigit(matrix[x][y]) {
		shouldCheckB = false
	}

	if x >= 0 && shouldCheckA {
		c := Coordinate{x, y}
		fromMap, present := mapOfPartNumbers[c]
		if present {
			gearRatio *= fromMap
			nAdjecentPartNumbers++
		}
	}

	x = row - 1
	y = col + 1
	if x >= 0 && y < n && shouldCheckB {
		c := Coordinate{x, y}
		fromMap, present := mapOfPartNumbers[c]
		if present {
			gearRatio *= fromMap
			nAdjecentPartNumbers++
		}
	}

	// check right cell
	x = row
	y = col + 1
	if y < n {
		c := Coordinate{x, y}
		fromMap, present := mapOfPartNumbers[c]
		if present {
			gearRatio *= fromMap
			nAdjecentPartNumbers++
		}
	}

	// check below row
	x = row + 1
	y = col + 1
	shouldCheckC := true
	if x < n && y < n {
		c := Coordinate{x, y}
		fromMap, present := mapOfPartNumbers[c]
		if present {
			shouldCheckC = false
			gearRatio *= fromMap
			nAdjecentPartNumbers++
		}
	}

	x = row + 1
	y = col
	shouldCheckD := true
	if x < n && unicode.IsDigit(matrix[x][y]) {
		shouldCheckD = false
	}

	x = row + 1
	y = col
	if x < n && shouldCheckC {
		c := Coordinate{x, y}
		fromMap, present := mapOfPartNumbers[c]
		if present {
			gearRatio *= fromMap
			nAdjecentPartNumbers++
		}
	}

	x = row + 1
	y = col - 1
	if x < n && y >= 0 && shouldCheckD {
		c := Coordinate{x, y}
		fromMap, present := mapOfPartNumbers[c]
		if present {
			gearRatio *= fromMap
			nAdjecentPartNumbers++
		}
	}

	if nAdjecentPartNumbers == 2 {
		return gearRatio
	}
	return 0
}

func getMapOfPartNumbers(s string) map[Coordinate]int {
	matrix := convertStringToMatrix(s)
	n := len(matrix)

	//key: Coordinate in matrix, value: the whole number in the matrix
	mapOfPartNumbers := make(map[Coordinate]int)

	for row := 0; row < n; row++ {
		for col := 0; col < n; {
			line := matrix[row]

			value := line[col]
			if !unicode.IsDigit(value) {
				col++
				continue
			}

			lastDigitIndex := col
			coordinatesOfCurrentNum := []Coordinate{}
			numberIsAPartNumber := false

			// traverse all the digits of current number
			for lastDigitIndex < n && unicode.IsDigit(line[lastDigitIndex]) {
				if isDigitAdjecentToSymbol(matrix, row, lastDigitIndex) {
					numberIsAPartNumber = true
				}

				coordinate := Coordinate{row, lastDigitIndex}
				coordinatesOfCurrentNum = append(coordinatesOfCurrentNum, coordinate)

				lastDigitIndex++
			}

			if numberIsAPartNumber {
				r := line[col:lastDigitIndex]
				num := fromRunesToNum(r)
				for _, v := range coordinatesOfCurrentNum {
					mapOfPartNumbers[v] = num
				}
			}

			col = lastDigitIndex + 1
		}
	}

	return mapOfPartNumbers
}

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	folder := "input_files"
	fileName := "input"
	fileExtension := ".txt"
	fileContent, err := os.ReadFile(folder + "/" + fileName + fileExtension)
	if err != nil {
		log.Fatal(err)
	}

	s := string(fileContent)
	s = strings.TrimSpace(s)

	// sum := getSumOfPartNumbers(s) // part 1
	sum := getSumOfGearRatios(s)
	fmt.Printf("sum of gear ratios is: %v ", sum)
}

func convertStringToMatrix(s string) [][]rune {
	var m [][]rune
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)

		lineAsRunes := []rune(line)
		m = append(m, lineAsRunes)
	}
	return m
}

func getSumOfPartNumbers(s string) int {
	partNumbers := getPartNumbers(s)
	sum := 0
	for _, v := range partNumbers {
		sum += v
	}
	return sum
}

func getPartNumbers(s string) []int {
	matrix := convertStringToMatrix(s)
	n := len(matrix)
	var partNumbers []int

	for row := 0; row < n; row++ {
		for col := 0; col < n; {
			line := matrix[row]

			// advance while we are not in a number
			if !unicode.IsDigit(line[col]) {
				col++
				continue
			}

			lastDigitIndex := col
			// a part number has any of its digits adjecent to a symbol
			numberIsAPartNumber := false
			// while we are in the same number, advance index
			for lastDigitIndex < n && unicode.IsDigit(line[lastDigitIndex]) {
				if isDigitAdjecentToSymbol(matrix, row, lastDigitIndex) {
					numberIsAPartNumber = true
				}
				lastDigitIndex++
			}

			if numberIsAPartNumber {
				// we can now construct the number as we know the start and end indexes
				r := line[col:lastDigitIndex]
				num := fromRunesToNum(r)
				partNumbers = append(partNumbers, num)
			}

			// we need to skip the current number
			col = lastDigitIndex + 1
		}
	}
	// fmt.Printf("partNumbers: %v", partNumbers)
	return partNumbers
}

// isDigitAdjecentToSymbol returns if matrix[row][col] is adjecent to a symbol.
func isDigitAdjecentToSymbol(matrix [][]rune, row, col int) bool {
	n := len(matrix)

	if col-1 >= 0 {
		v := matrix[row][col-1]
		if !unicode.IsDigit(v) && string(v) != "." {
			return true
		}
	}

	if row-1 >= 0 && col-1 >= 0 {
		v := matrix[row-1][col-1]
		if !unicode.IsDigit(v) && string(v) != "." {
			return true
		}
	}

	if row-1 >= 0 {
		v := matrix[row-1][col]
		if !unicode.IsDigit(v) && string(v) != "." {
			return true
		}
	}

	if row-1 >= 0 && col+1 < n {
		v := matrix[row-1][col+1]
		if !unicode.IsDigit(v) && string(v) != "." {
			return true
		}
	}

	if col+1 < n {
		v := matrix[row][col+1]
		if !unicode.IsDigit(v) && string(v) != "." {
			return true
		}
	}

	if row+1 < n && col+1 < n {
		v := matrix[row+1][col+1]
		if !unicode.IsDigit(v) && string(v) != "." {
			return true
		}
	}

	if row+1 < n {
		v := matrix[row+1][col]
		if !unicode.IsDigit(v) && string(v) != "." {
			return true
		}
	}

	if row+1 < n && col-1 >= 0 {
		v := matrix[row+1][col-1]
		if !unicode.IsDigit(v) && string(v) != "." {
			return true
		}
	}

	return false
}

func fromRunesToNum(runesOfNum []rune) int {
	wholeNumStr := string(runesOfNum)
	wholeNumInt, err := strconv.Atoi(wholeNumStr)
	if err != nil {
		log.Fatal(err)
	}
	return wholeNumInt
}

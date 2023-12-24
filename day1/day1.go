package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	folder := "input_files"
	fileName := "input.txt"
	fileContent, err := os.ReadFile(folder + "/" + fileName)
	if err != nil {
		log.Fatal(err)
	}

	s := string(fileContent)
	s = strings.TrimSpace(s)
	lines := getLinesFromFileContent(s)

	// sum, err := GetSumOfCalibrationValues(lines) // Part 1
	sum, err := GetSumOfCalibrationValuesPart2(lines)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("result is %v: ", sum)
}

func getLinesFromFileContent(s string) []string {
	return strings.Split(s, "\n")
}

func GetSumOfCalibrationValues(lines []string) (int, error) {
	sum := 0
	for _, line := range lines {
		calibrationValueStr, err := getCalibrationValue(line)
		if err != nil {
			return 0, err
		}

		calibrationValue, err := strconv.Atoi(calibrationValueStr)
		if err != nil {
			return 0, err
		}

		sum += calibrationValue
	}

	return sum, nil
}

func getCalibrationValue(input string) (string, error) {
	firstDigit := firstDigitFromString(input)

	resversed, err := reverseString(input)
	if err != nil {
		return "", err
	}

	lastDigit := firstDigitFromString(resversed)

	return firstDigit + lastDigit, nil
}

func reverseString(input string) (string, error) {
	runes := []rune(input)
	slices.Reverse(runes)
	return string(runes), nil
}

func firstDigitFromString(input string) string {
	for _, c := range input {
		if unicode.IsDigit(c) {
			return string(c)
		}
	}
	return ""
}

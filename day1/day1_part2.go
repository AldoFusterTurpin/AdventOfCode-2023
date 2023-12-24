package main

import (
	"strconv"
	"unicode"
)

func spelledNumbers() []string {
	return []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
}

func spelledNumbersReversed() []string {
	return []string{"eno", "owt", "eerht", "ruof", "evif", "xis", "neves", "thgie", "enin"}
}

func valueFromSpelledNumbers(number string) string {
	m := map[string]string{
		"one": "1", "two": "2", "three": "3", "four": "4",
		"five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9",

		"eno": "1", "owt": "2", "eerht": "3", "ruof": "4",
		"evif": "5", "xis": "6", "neves": "7", "thgie": "8", "enin": "9",
	}

	return m[number]
}

func GetSumOfCalibrationValuesPart2(lines []string) (int, error) {
	sum := 0
	for _, line := range lines {
		calibrationValueStr, err := getCalibrationValuePart2(line)
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

func getCalibrationValuePart2(input string) (string, error) {
	firstDigit := firstDigitFromStringPart2(input, spelledNumbers())

	resversed, err := reverseString(input)
	if err != nil {
		return "", err
	}

	lastDigit := firstDigitFromStringPart2(resversed, spelledNumbersReversed())

	return firstDigit + lastDigit, nil
}

func firstDigitFromStringPart2(input string, digitsSpelledOutWithLetters []string) string {
	for i, c := range input {
		if unicode.IsDigit(c) {
			return string(c)
		}
		if v, b := getSpelledDigitAtBeginningOfS(input[i:], digitsSpelledOutWithLetters); b {
			return v
		}
	}
	return ""
}

// getSpelledDigitAtBeginningOfS will iterate over digitsSpelledOutWithLetters to check if matches
// any spelled digit and return the matched number and true. Otherwise will neturn "" and false.
// Example:
// s -> "two1nine"
// matches "one" ? -> no match
// matches "two" -> match -> return 2
func getSpelledDigitAtBeginningOfS(s string, digitsSpelledOutWithLetters []string) (string, bool) {
	for _, digitSpelledOutWithLetters := range digitsSpelledOutWithLetters {
		n := len(digitSpelledOutWithLetters)
		if n > len(s) {
			continue
		}
		tryNumber := s[:n]
		if tryNumber == digitSpelledOutWithLetters {
			return valueFromSpelledNumbers(digitSpelledOutWithLetters), true
		}
	}
	return "", false
}

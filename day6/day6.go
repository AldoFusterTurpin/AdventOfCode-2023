package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	folder := "input_files"
	fileName := "input"
	fileExtension := ".txt"

	fileContentBytes, err := os.ReadFile(folder + "/" + fileName + fileExtension)
	if err != nil {
		log.Fatal(err)
	}

	fileContent := string(fileContentBytes)
	partTwo := true

	times, distances := processFileContent(fileContent, partTwo)
	// fmt.Println(times)
	// fmt.Println(distances)

	result := getMultiplicationOfAllTheNumberOfWaysToBeatTheRecords(times, distances)
	fmt.Println("result is", result) // result is 30077773
}

func getMultiplicationOfAllTheNumberOfWaysToBeatTheRecords(times, distances []int) int {
	if len(times) != len(distances) {
		log.Fatal("error: number of times is different than number of distances.")
	}

	result := 1
	for i := 0; i < len(times); i++ {
		numberOfWaysWeCanBeatTheRecord := getNumberOfWaysWeCanBeatTheRecord(times, i, distances)
		if numberOfWaysWeCanBeatTheRecord > 0 {
			result *= numberOfWaysWeCanBeatTheRecord
		}
	}
	return result
}

func getNumberOfWaysWeCanBeatTheRecord(times []int, i int, distances []int) int {
	time := times[i]
	recordDistance := distances[i]
	numberOfWaysWeCanBeatTheRecord := 0

	for holdButtonForNMilliseconds := 0; holdButtonForNMilliseconds < time; holdButtonForNMilliseconds++ {
		millisecondsItWillBeMoving := time - holdButtonForNMilliseconds
		totalDistance := millisecondsItWillBeMoving * holdButtonForNMilliseconds
		if totalDistance > recordDistance {
			numberOfWaysWeCanBeatTheRecord++
		}
	}
	return numberOfWaysWeCanBeatTheRecord
}

func processFileContent(fileContent string, partTwo bool) (times, distances []int) {
	replaceWith := " "
	if partTwo {
		replaceWith = ""
	}

	lines := strings.Split(fileContent, "\n")

	firstLine := lines[0]
	multipleSpaces := regexp.MustCompile(`\s+`)

	firstLine = multipleSpaces.ReplaceAllString(firstLine, replaceWith)
	firstLineSlice := strings.Split(firstLine, ":")

	timesStr := strings.TrimSpace(firstLineSlice[1])

	timesStrSlice := strings.Split(timesStr, " ")
	times = fromSliceOfStringsToSliceOfInts(timesStrSlice)

	secondLine := lines[1]
	secondLine = multipleSpaces.ReplaceAllString(secondLine, replaceWith)

	secondLineSlice := strings.Split(secondLine, ":")

	distanceStr := strings.TrimSpace(secondLineSlice[1])

	distanceStrSlice := strings.Split(distanceStr, " ")
	distances = fromSliceOfStringsToSliceOfInts(distanceStrSlice)

	return times, distances
}

func fromSliceOfStringsToSliceOfInts(s []string) []int {
	res := make([]int, len(s))
	for i, x := range s {
		number, err := strconv.Atoi(x)
		if err != nil {
			log.Fatal(number)
		}
		res[i] = number
	}
	return res
}

package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// desiredNumberOfRangePairs contains the desired number of range pairs we want our input to be broken into.
// It is used to break the ranges into samller ones as we will execute each range in a different Go routine to
// sepped up the whole process.
//
// The results are really impressive, we can process the "input.txt" in a little bit more than
// one minute, and it contains really big ranges. When working in Go, 1 OS Thread can be handling many Go gorutines
// (so the mapping is not 1 Thread -> 1 goroutine, it is 1 Thread -> n goroutines),
// This is something the Go runtime is responsible of.
// I simply love Go, for its concurrency model and for many other reasons.
//
// If desiredNumberOfRangePairs is less than 0, we will not break the ranges into smaller ones (when the ranges are small, we
// don't want to break them into smaller ones as it is not useful, like in the "sample.txt").
// For "input.txt" a good number is something close to 10_000 (check "notes.txt")
// Ideally we could decide this variable at runtime, but this is a simplification.
const desiredNumberOfRangePairs = 10_000

var (
	multipleSpacesRegex = regexp.MustCompile(`\s+`)
)

func main() {
	folder := "input_files"
	fileName := "input"
	fileExtension := ".txt"
	fileContent, err := os.ReadFile(folder + "/" + fileName + fileExtension)
	if err != nil {
		log.Fatal(err)
	}

	sum := GetLowestLocationOfSeedPairsConcurrent(string(fileContent), desiredNumberOfRangePairs)
	fmt.Printf("lowest location of range of seed pairs is %v\n", sum)
}

func GetLowestLocationOfAllSeeds(s string) int {
	almanac := convertStringToAlmanac(s)
	return almanac.getLowestLocationOfAllSeeds()
}

type MapFromSourceToDestination struct {
	from        string // e.g: seed, soil
	destination string // e.g: soil, fertilizer
	ranges      []Range
}

func (m MapFromSourceToDestination) GetDestinationValue(source int) int {
	for _, r := range m.ranges {
		if source >= r.sourceRangeStart && source <= r.sourceRangeStart+r.rangeLength-1 {
			diff := r.sourceRangeStart - r.destinationRangeStart
			offset := source - r.sourceRangeStart
			return r.sourceRangeStart + offset - diff
		}
	}
	// Any source numbers that aren't mapped correspond to the same destination number.
	return source
}

type Range struct {
	destinationRangeStart, sourceRangeStart, rangeLength int
}

type Almanac struct {
	seedsToBePlanted []int
	maps             []MapFromSourceToDestination
}

func (a Almanac) getLowestLocationOfAllSeeds() int {
	var locations []int

	for _, seed := range a.seedsToBePlanted {
		source := seed
		for _, m := range a.maps {
			source = m.GetDestinationValue(source)
		}
		locations = append(locations, source)
	}
	return getMinFromSlice(locations)
}

func getMinFromSlice(locations []int) int {
	first := true
	min := 0

	for _, v := range locations {
		if first {
			first = false
			min = v
		} else if v < min {
			min = v
		}
	}

	return min
}

func convertStringToAlmanac(s string) Almanac {
	almanac := Almanac{}

	s = strings.TrimSpace(s)
	lines := strings.Split(s, "\n\n")
	seedsStr := lines[0]
	seedsStr = multipleSpacesRegex.ReplaceAllString(seedsStr, " ")
	seedsStrSplitted := strings.Split(seedsStr, ":")

	seedsToBePlantedStr := seedsStrSplitted[1]
	almanac.seedsToBePlanted = convertStringToSliceOfInts(seedsToBePlantedStr)
	for _, line := range lines[1:] {
		runes := []rune(line)
		r := runes[0]
		if unicode.IsLetter(r) {
			m := MapFromSourceToDestination{}

			lineSplitted := strings.Split(line, "\n")
			lineZero := lineSplitted[0]
			lineZeroSplitted := strings.Split(lineZero, "-")

			m.from = lineZeroSplitted[0]
			lastPartOfLineSplitted := strings.Split(lineZeroSplitted[2], " ")
			m.destination = lastPartOfLineSplitted[0]

			for _, ll := range lineSplitted[1:] {
				ints := convertStringToSliceOfInts(ll)
				r := Range{
					destinationRangeStart: ints[0],
					sourceRangeStart:      ints[1],
					rangeLength:           ints[2],
				}
				m.ranges = append(m.ranges, r)
			}
			almanac.maps = append(almanac.maps, m)
		}
	}
	return almanac
}

func convertStringToSliceOfInts(numbers string) []int {
	numbers = strings.TrimSpace(numbers)
	var result []int

	numbersSLiceStr := strings.Split(numbers, " ")
	for _, s := range numbersSLiceStr {
		v, err := strconv.Atoi(s)
		if err != nil {
			// Unexpected error, just for simplification...
			log.Fatal(err)
		}
		result = append(result, v)
	}
	return result
}

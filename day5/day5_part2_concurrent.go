package main

import (
	"log"
	"sort"
	"sync"
)

type ListOfRangePairs struct {
	rangePairs []RangePair
}

type RangePair struct {
	id           int
	startOfRange int
	length       int
}

func GetLowestLocationOfSeedPairsConcurrent(s string, desiredNumberOfRangePairs int) int {
	almanac := convertStringToAlmanac(s)
	return almanac.getLowestLocationOfSeedPairsConcurrent(desiredNumberOfRangePairs)
}

func createSubRange(rangePair RangePair) []RangePair {
	divideFactor := 2
	n := rangePair.length / divideFactor

	if rangePair.length%divideFactor == 0 {
		rp1 := RangePair{
			startOfRange: rangePair.startOfRange,
			length:       n,
		}

		rp2 := RangePair{
			startOfRange: rangePair.startOfRange + n,
			length:       n,
		}

		return []RangePair{rp1, rp2}
	}

	rp1 := RangePair{
		startOfRange: rangePair.startOfRange,
		length:       n,
	}

	rp2 := RangePair{
		startOfRange: rangePair.startOfRange + n,
		length:       rangePair.length - n,
	}

	return []RangePair{rp1, rp2}
}

func (a Almanac) getLowestLocationOfSeedPairsConcurrent(desiredNumberOfRangePairs int) int {
	listOfRangepairs := fromSliceOfIntsToRangePairs(a.seedsToBePlanted)
	listOfRangepairs = breakRangesIntoSmallerOnesMultipleTimes(listOfRangepairs, desiredNumberOfRangePairs)

	var wg sync.WaitGroup
	resChannel := make(chan int)

	for _, rangePair := range listOfRangepairs.rangePairs {
		// process each rangePair in a go routine
		wg.Add(1)
		go sendMinFromRangePairToResChannel(rangePair, a.maps, resChannel, &wg)
	}

	// Why we need to wait in a different go routine?
	// because we need the main goroutine reach the code where
	// we iterate over the result channel. If we place the waitgroup on the main goroutine (without the "go" keyword),
	// the code will never reach the "close(resChannel)" because the thing that is controlling the close of the channel is in fact
	// the waitgroup, so we will end up with we waiting using the waitgroup to close the channel + we never iterating the res channel
	// as that iteration happens after we wait with the waitgroup.
	// With the wait in a different go routine we are ensuring:
	// - in the main goroutine, iterate over the result channel until someone closes the resChannel
	// - in a different go routine, wait until we have sent all the values and then "signal" that closing the channel ->
	// that will make getMinFromChannel stop iterating over the resChannel, providing the final result.
	go func() {
		wg.Wait()

		// it is safe to close the channel as all the workers have finished (because we have reached the wg.Wait()).
		// If we don't close it here, we should use a counter of how many workers are still writing, the mutex
		// already does that.Simpler.
		close(resChannel)
	}()

	// main goroutine gets blocked in getMinFromChannel when iterating over the channel until someone closes it
	return getMinFromChannel(resChannel)
}

func breakRangesIntoSmallerOnesMultipleTimes(listOfRangepairs ListOfRangePairs, desiredNumberOfRangePairs int) ListOfRangePairs {
	// this means we don't want to break the ranges into smaller ones.
	if desiredNumberOfRangePairs < 0 {
		return listOfRangepairs
	}

	n := len(listOfRangepairs.rangePairs)

	for n < desiredNumberOfRangePairs {
		listOfRangepairs = breakRangesIntoSmallerOnes(listOfRangepairs)
		n = len(listOfRangepairs.rangePairs)
	}

	return listOfRangepairs
}

// breakRangesIntoSmallerOnes takes a list of range pairs as input and returns another
// list of range pairs but with more ranges that are smaller in order to leverage
// concurrency treating each range in a go routine.
func breakRangesIntoSmallerOnes(listOfRangepairs ListOfRangePairs) ListOfRangePairs {
	var rangePairs []RangePair
	for _, rangePair := range listOfRangepairs.rangePairs {
		rangePairs = append(rangePairs, createSubRange(rangePair)...)
	}
	return ListOfRangePairs{rangePairs: rangePairs}
}

// getMinFromChannel iterates over the resChannel and returns the minimum value of it.
func getMinFromChannel(resChannel <-chan int) int {
	min := 0
	first := true
	for x := range resChannel {
		if first {
			min = x
			first = false
		} else if x < min {
			min = x
		}
	}

	return min
}

// sendMinFromRangePairToResChannel gets the minimmum from the range pair using the maps and sends the result to
// the resChannel.
func sendMinFromRangePairToResChannel(rangePair RangePair, maps []MapFromSourceToDestination, resChannel chan<- int, wg *sync.WaitGroup) {
	// fmt.Printf("getMinFromRange() with rangePair: %v\n", rangePair)

	defer wg.Done()
	min := 0
	first := true

	x := rangePair.startOfRange

	// length-1 because a range with length 1 means just one number:
	// Range 1(start), 1(length) means [1]
	// Range 1, 2 means [1, 2]
	// Range 1, 3 means [1, 2, 3]
	for i := 0; i < rangePair.length-1; i++ {
		res := iterateOverTheMaps(x, maps)

		if first {
			min = res
			first = false
		} else if res < min {
			min = res
		}

		x++
	}
	resChannel <- min
}

func iterateOverTheMaps(x int, maps []MapFromSourceToDestination) int {
	for _, m := range maps {
		x = m.GetDestinationValue(x)
	}
	return x
}

func fromSliceOfIntsToRangePairs(slice []int) ListOfRangePairs {
	n := len(slice)
	if n%2 != 0 {
		log.Fatal("odd number of numbers, can not create range pairs")
	}

	rangeId := 0
	var rangePairs []RangePair
	for i := 0; i < n; i += 2 {
		rp := RangePair{
			id:           rangeId,
			startOfRange: slice[i],
			length:       slice[i+1],
		}
		rangePairs = append(rangePairs, rp)
		rangeId++
	}

	sort.Slice(rangePairs, func(i, j int) bool {
		return rangePairs[i].startOfRange < rangePairs[j].startOfRange
	})

	listOfRangePairs := ListOfRangePairs{
		rangePairs: rangePairs,
	}

	return listOfRangePairs
}

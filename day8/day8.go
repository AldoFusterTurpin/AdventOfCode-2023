package main

import (
	"fmt"
	"log"
	"os"
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

	inputString := string(fileContentBytes)
	inputString = strings.TrimSpace(inputString)
	splitStr := strings.Split(inputString, "\n\n")
	instructionsStr := splitStr[0]
	nodesStr := splitStr[1]

	// fmt.Printf("instructions is:\n%v\n\n", instructionsStr)
	// fmt.Printf("nodes are:\n%v", nodesStr)

	nodes := convertStringToNodes(nodesStr)
	// fmt.Printf("nodes are:\n%+v", nodes)

	mapOfNodes := buildMapOfNodes(nodes)
	// fmt.Printf("mapOfNodes is:\n%+v", mapOfNodes)

	nSteps := getNecessaryStepsToReachDestination(instructionsStr, mapOfNodes)
	fmt.Printf("%v steps required to reach destination", nSteps)
}

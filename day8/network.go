package main

import (
	"strings"
)

type Node struct {
	Id           string
	LeftElement  string
	RightElement string
}

func convertStringToNodes(inputString string) []Node {
	var nodes []Node

	lines := strings.Split(inputString, "\n")
	for _, line := range lines {
		// AAA = (BBB, CCC)
		line = strings.ReplaceAll(line, " ", "")
		line = strings.ReplaceAll(line, "(", "")
		line = strings.ReplaceAll(line, ")", "")
		// AAA=BBB,CCC

		splitStr := strings.Split(line, "=")
		// []string { "AAA", "BBB,CCC"}

		id := splitStr[0] // "AAA"

		nodeElements := splitStr[1] // "BBB,CCC"
		splitStr = strings.Split(nodeElements, ",")
		leftElement := splitStr[0]  // BBB
		rightElement := splitStr[1] // CCC

		node := Node{
			Id:           id,
			LeftElement:  leftElement,
			RightElement: rightElement,
		}
		nodes = append(nodes, node)
	}
	return nodes
}

// buildMapOfNodes returns a map where key is Node ID and value the node itself.
func buildMapOfNodes(nodes []Node) map[string]Node {
	mapOfNodes := make(map[string]Node)
	for _, node := range nodes {
		mapOfNodes[node.Id] = node
	}
	return mapOfNodes
}

func getNecessaryStepsToReachDestination(instructionsStr string, mapOfNodes map[string]Node) int {
	startNodeId := "AAA"
	endNodeId := "ZZZ"
	nSteps := 0
	currentNode := mapOfNodes[startNodeId]

	currentInstructionIndex := 0
	for {
		for currentInstructionIndex = 0; currentInstructionIndex < len(instructionsStr); currentInstructionIndex++ {
			instruction := instructionsStr[currentInstructionIndex]
			nSteps++

			if instruction == 'L' {
				currentNode = mapOfNodes[currentNode.LeftElement]
			} else if instruction == 'R' {
				currentNode = mapOfNodes[currentNode.RightElement]
			}

			if currentNode.Id == endNodeId {
				return nSteps
			}
		}
		// once we have consumed the full instruction string, we start iterating again from the beginning
		currentInstructionIndex = 0
	}
}

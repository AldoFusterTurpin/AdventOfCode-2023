package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	folder := "input_files"
	fileName := "sample"
	fileExtension := ".txt"

	fileContentBytes, err := os.ReadFile(folder + "/" + fileName + fileExtension)
	if err != nil {
		log.Fatal(err)
	}

	fileContent := string(fileContentBytes)

	handsWithBids, err := getHandsWithBidsFromFileContent(fileContent)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("handsWithBids: %+v", handsWithBids)
	// fmt.Println("result is", result)
}

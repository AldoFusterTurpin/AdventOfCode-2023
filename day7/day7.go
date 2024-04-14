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

	fileContent := string(fileContentBytes)
	fileContent = strings.TrimSpace(fileContent)

	hs, err := getHandsWithBidsFromFileContent(fileContent)
	if err != nil {
		log.Fatal(err)
	}
	result := GetTotalWinnings(hs)
	// fmt.Printf("hs: %+v", hs)
	fmt.Println("result is", result)
}

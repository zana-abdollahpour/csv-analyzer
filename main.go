package main

import (
	"fmt"
	"log"

	"csvanalyzer/pkg"
)

func main() {
	pkg.Greet()

	searchedTerm := pkg.GetSearchTerm()

	inputFileNames, inputNamesErr := pkg.GetInputFileNames()
	if inputNamesErr != nil {
		log.Fatal("unexpected file open error: ", inputNamesErr)
	}

	for _, fileName := range inputFileNames {
		pkg.SaveMatchingEntries(fileName, searchedTerm)
	}

	fmt.Println("---> Analyze completed! <---")
}

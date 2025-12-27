package main

import (
	"fmt"
	"log"

	"csvanalyzer/pkg"
)

func main() {
	pkg.Greet()

	var searchedTerm string
	for {
		searchedTerm = pkg.GetSearchTerm()
		if searchedTerm == "" {
			fmt.Println("Please enter a non-empty string!")
			continue
		} else {
			break
		}

	}

	inputFileNames, inputNamesErr := pkg.GetInputFileNames()
	if inputNamesErr != nil {
		log.Fatal("unexpected file open error: ", inputNamesErr)
	}

	for _, fileName := range inputFileNames {
		fmt.Printf("--> Started Analysis: %s <--", fileName)
		pkg.SaveMatchingEntries(fileName, searchedTerm)
		fmt.Printf("--> Finished Analysis: %s <--", fileName)
	}

	fmt.Println("---> Analysis completed! <---")

	pkg.ShutdownSystemWithDelay(0)
}

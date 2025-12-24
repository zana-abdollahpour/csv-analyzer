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
		fmt.Println("-------------------------------------------------------------------> Started Analysis")
		pkg.SaveMatchingEntries(fileName, searchedTerm)
		fmt.Println("-------------------------------------------------------------------> Finished Analysis")
	}

	fmt.Println("---> Analysis completed! <---")

	pkg.ShutdownSystemWithDelay(0)
}

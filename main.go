package main

import (
	"fmt"
	"log"
	"strconv"

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

	finishCallback, cbErr := pkg.AskAfterFinishBehavior()
	if cbErr != nil {
		fmt.Println(cbErr)
	}
	if finishCallback != nil {
		delay, err := strconv.Atoi(pkg.GetOperationDelay())
		if err != nil {
			fmt.Println("Invalid delay value, the default of 10 will be applied")
			delay = 10
		}
		defer finishCallback(delay)
	}

	inputFileNames, inputNamesErr := pkg.GetInputFileNames()
	if inputNamesErr != nil {
		log.Fatal("unexpected file open error: ", inputNamesErr)
	}

	for _, fileName := range inputFileNames {
		fmt.Printf("--> Started Analysis: %s <--\n", fileName)
		pkg.SaveMatchingEntries(fileName, searchedTerm)
		fmt.Printf("--> Finished Analysis: %s <--\n", fileName)
	}

	fmt.Println("---> Analysis completed! <---")
}

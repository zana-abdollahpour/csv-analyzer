package pkg

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func GetInputFileNames() ([]string, error) {
	entries, readDirErr := os.ReadDir("./input")
	if readDirErr != nil {
		return nil, readDirErr
	}

	fileNames := make([]string, 0)
	for _, entry := range entries {
		info, infoErr := entry.Info()
		if infoErr != nil {
			return nil, infoErr

		}

		if strings.HasSuffix(info.Name(), ".csv") {
			fileNames = append(fileNames, info.Name())
		}
	}

	return fileNames, nil
}

func SaveMatchingEntries(fileName, searchedTerm string) {
	const (
		fileWriteFlags      = os.O_CREATE | os.O_WRONLY | os.O_TRUNC
		fileWritePermission = os.FileMode(0644)
		dirPermission       = os.FileMode(0755)
	)

	inputFilePath := fmt.Sprintf("./input/%s", fileName)
	sourceFile, sourceOpenErr := os.Open(inputFilePath)
	if sourceOpenErr != nil {
		fmt.Println("!!! error opening source file: ", sourceOpenErr)
		return
	}
	defer sourceFile.Close()

	reader := csv.NewReader(sourceFile)

	outputDir := "./output/" + searchedTerm
	if outputDrr := os.MkdirAll(outputDir, dirPermission); outputDrr != nil {
		fmt.Println("!!! error creating output directory:", outputDrr)
		return
	}

	outputFilePath := outputDir + "/" + fileName
	outputFile, outputOpenErr := os.OpenFile(outputFilePath, fileWriteFlags, fileWritePermission)
	if outputOpenErr != nil {
		fmt.Println("!!! error opening output file: ", outputOpenErr)
		return
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	columns, columnsErr := reader.Read()
	if columnsErr != nil {
		fmt.Println("Error reading line:", columnsErr)
		return
	}
	writer.Write(columns)

	var readRecordCount = 1
	for {
		fmt.Printf("\r%d", readRecordCount)
		record, err := reader.Read()
		readRecordCount++
		if err == io.EOF {
			continue
		}

		if err != nil {
			continue
		}

		for _, val := range record {
			if strings.Contains(val, searchedTerm) {
				if err := writer.Write(record); err != nil {
					fmt.Println("Error writing line:", err)
					break
				}
			}
		}
	}

}

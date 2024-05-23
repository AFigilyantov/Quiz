package main

import (
	"encoding/csv"
	"log"
	"os"
)

func readCSVFile(filePath string) [][]string {

	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal("Cann't open the file"+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Cann't parse file"+filePath, err)
	}
	return records
}

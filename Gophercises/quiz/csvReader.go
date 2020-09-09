package main

import (
	"encoding/csv"
	"log"
	"os"
)

func readCsvLineWise(filename string) [][]string{
	lines, err := csvReader(filename)
	if err != nil {
		log.Fatal(err)
		return [][]string{}
	}
	return lines
}

func csvReader(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}


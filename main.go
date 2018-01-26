package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func ParseCsv(filename string) ([][]string, error) {
	csvFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()

	r := csv.NewReader(bufio.NewReader(csvFile))

	lines := [][]string{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		lines = append(lines, record)
	}

	return lines, nil
}

func main() {
	fmt.Println("Starting CTT backfill")

	records, _ := ParseCsv("/Users/eris.mai/Downloads/search-results-2018-01-25T19_20_19.070-0800.csv")
	fmt.Printf("Parsed [%v] records from csv file", len(records))

	m := make(map[string]int)

	for i := 1; i < len(records); i++ {
		m[records[i][6]]++
	}
	fmt.Println("*********")
	for k, v := range m {
		fmt.Println(k, ": ", v)
	}
}

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Started - A primtive CSV cleaner written in GOLang!")

	if len(os.Args) < 2 {
		log.Fatal("Usage : ./program [file_path]")
	}
	file_path := os.Args[1]

	records := readCsv(file_path)

	cleanWhiteSpace(records)

	fmt.Println(records)

	fmt.Println("Exiting program!")
}

func readCsv(file_path string) [][]string {
	file, err := os.Open(file_path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	r := csv.NewReader(file)

	var records [][]string

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		records = append(records, record)
	}
	return records
}

func cleanWhiteSpace(records [][]string) [][]string {

	for i := range records {
		for j := range records[i] {
			records[i][j] = strings.TrimSpace(records[i][j])
		}
	}
	return records
}

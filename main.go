package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Started - A primtive CSV cleaner written in GOLang!")

	if len(os.Args) < 2 {
		log.Fatal("Usage : ./program [file_path]")
	}

	file_path := os.Args[1]

	records := readCsv(file_path)

	cleanWhiteSpace(records)
	handleNullValues(records)
	normalizeDates(records)
	

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

func handleNullValues(records [][]string)[][]string {
	
	for i := range records { 
		for j := range records [i] {
				
			switch strings.ToLower(records[i][j]) {
			case "n/a", "none", "null", "-", "na":
    	records[i][j] = ""
		}
	}
}
	return records
}

func normalizeDates(records[][]string)[][]string{
	
	layouts := []string{"02/01/2006", "2006/01/02", "2006-01-02"}

	for i := range records{
		for j := range records [i]{
			for _, layout := range layouts{
				t, err := time.Parse(layout, records[i][j])
				if err == nil{
					records[i][j] = t.Format("2006-01-02")
					break
				}
			}
			
		
		}
	}
	return records
}


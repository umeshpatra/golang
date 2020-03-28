package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func SaveCSVFile(rows [][]string) {
	file, err := os.Create("customer.csv")

	if err != nil {
		log.Fatalln(err)
	}

	csvwriter := csv.NewWriter(file)

	for _, row := range rows {
		_ = csvwriter.Write(row)
	}
	csvwriter.Flush()
	file.Close()
}

func LoadCSVFile(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var texts []string
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		texts = append(texts, scanner.Text())
	}
	fmt.Println(texts)
}
func main() {

	rows := [][]string{
		{"Name", "City", "Language"},
		{"Subrat", "IRELAND", "UK-ENG"},
		{"Ashok", "SEATTLE", "US-ENG"},
		{"Sameer", "ODISHA", "IN-HINDI"},
	}
	SaveCSVFile(rows)

	LoadCSVFile("customer.csv")

}

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/xuri/excelize/v2"
)

var (
	FILE_NAME        = "./metrics/" + os.Getenv("FILE_NAME") + ".xlsx"
	SPREADSHEET_NAME = os.Getenv("SPREADSHEET")
	CSV_NAME         = "metrics/Metrics.csv"
)

func main() {
	go updateCsv()

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, CSV_NAME)
	})
	http.ListenAndServe(":8081", nil)
}

func updateCsv() {
	fmt.Println("Updating .csv file every one second.")
	for {
		rows, err := getXlxsRows(FILE_NAME, SPREADSHEET_NAME)
		if err != nil {
			fmt.Println(err)
		}

		err = writeCsv(rows)
		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(1 * time.Second)
	}
}

// writeCsv Writes down all received data to an .CSV file
func writeCsv(data [][]string) error {
	csvFile, err := os.Create(CSV_NAME)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
		return err
	}
	defer csvFile.Close()

	w := csv.NewWriter(csvFile)
	defer w.Flush()

	err = w.WriteAll(data)
	if err != nil {
		log.Fatalf("%s", err)
		return err
	}

	return nil
}

// getXlxsRows Opens an .xlsx file and reads returns all rows from it
func getXlxsRows(fileName, sheetName string) ([][]string, error) {
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		fmt.Printf("Error in opening file %s\n", err)
		return nil, err
	}
	defer closeSpreadsheet(f)

	rows, err := f.GetRows(sheetName)
	if err != nil {
		fmt.Printf("Wasn't able to get rows from spreadsheet %s from file %s\n%s\n", sheetName, fileName, err)
		return nil, err
	}

	return rows, nil
}

// closeSpreadsheet Closes the spreadsheet
func closeSpreadsheet(f *excelize.File) {
	if err := f.Close(); err != nil {
		fmt.Println(err)
	}
}

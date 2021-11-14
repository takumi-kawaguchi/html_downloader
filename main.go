package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	var outputFolderPath string
	print("Select a folder in which you download html files: ")
	fmt.Scan(&outputFolderPath)
	if !Exists(outputFolderPath) {
		os.Exit(1)
	}

	var downloadTargetCsvPath string
	print("Select a CSV file: ")
	fmt.Scan(&downloadTargetCsvPath)
	if !Exists(downloadTargetCsvPath) {
		fmt.Println("指定されたCSVは存在しません")
		os.Exit(1)
	}

	outputFolderPath = filepath.Join(outputFolderPath, time.Now().Format("20060102150405")+"_htmldownloder")
	err := os.Mkdir(outputFolderPath, 0777)
	if err != nil {
		log.Fatal(err)
	}

	csvFile, err := os.Open(downloadTargetCsvPath)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	csvContent := csv.NewReader(csvFile)
	records, err := csvContent.ReadAll()
	if err != nil {
		log.Fatal()
	}

	fmt.Println("Start requests...")
	for _, record := range records {
		// リクエスト
		res, err := http.Get(record[1])
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			log.Fatalf("statuscode: %v\n", res.StatusCode)
		}

		// 格納フォルダ作成
		folders := strings.Split(record[2], "\\")
		eachRequestTargetPath := outputFolderPath
		for _, folder := range folders {
			eachRequestTargetPath = filepath.Join(eachRequestTargetPath, folder)
			if !Exists(eachRequestTargetPath) {
				err := os.Mkdir(eachRequestTargetPath, 0777)
				if err != nil {
					log.Fatal(err)
				}
			}
		}

		// レスポンスボディをHTMLに詰める
		resHtml, err := os.Create(filepath.Join(eachRequestTargetPath, fmt.Sprintf("%s.html", record[0])))
		if err != nil {
			log.Fatal(err)
		}
		defer resHtml.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		resHtml.Write(body)
	}

	fmt.Println("All requests completed, htmls stored in your folders.")
}

func Exists(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}

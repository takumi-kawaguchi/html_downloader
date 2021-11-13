package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	var outputFolderPath string
	print("HTMLのダウンロード先フォルダを選択してください: ")
	fmt.Scan(&outputFolderPath)
	if !Exists(outputFolderPath) {
		os.Exit(1)
	}

	var downloadTargetCsvPath string
	print("ダウンロードHTMLをまとめたCSVを選択してください: ")
	fmt.Scan(&downloadTargetCsvPath)
	if !Exists(downloadTargetCsvPath) {
		os.Exit(1)
	}

	var rootFolderName = time.Now().Format("20060102150405") + "_htmldownloder"

	of, err := os.Create(outputFolderPath + rootFolderName)
	if err != nil {
		log.Fatal(err)
	}
	defer of.Close()

	f, err := os.Open(downloadTargetCsvPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal()
	}

	for _, record := range records {
		res, err := http.Get(record[1])
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			log.Fatalf("statuscode: %v\n", res.StatusCode)
		}

		f, err := os.Create(fmt.Sprintf(outputFolderPath+"%s.html", record[0]))
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		f.Write(body)
	}
}

func Exists(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}

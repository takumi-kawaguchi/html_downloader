package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	f, err := os.Open("access.csv")
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

		f, err := os.Create(fmt.Sprintf("output/%s.html", record[0]))
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

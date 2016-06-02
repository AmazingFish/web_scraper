package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// get meaning from weblio
func get_meaning(word string) string {
	var meanings []string
	time.Sleep(1000 * time.Millisecond)
	doc, err := goquery.NewDocument("http://ejje.weblio.jp/content/" + word)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".level0").Each(func(i int, s *goquery.Selection) {
		meanings = append(meanings, s.Text())
	})

	meaning := strings.Join(meanings, ",")

	return meaning
}

func main() {
	var rfp *os.File
	var wfp *os.File
	var err error

	if len(os.Args) < 2 {
		rfp = os.Stdin
	} else {
		rfp, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer rfp.Close()
	}

	if len(os.Args) < 3 {
		wfp = os.Stdin
	} else {
		wfp, err = os.OpenFile(os.Args[2], os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}
		defer wfp.Close()
	}

	var records [][]string

	scanner := bufio.NewScanner(rfp)
	for scanner.Scan() {
		text := scanner.Text()
		meaning := get_meaning(text)
		records = append(records, []string{meaning})
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	w := csv.NewWriter(wfp)
	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write buffered data to the underlying writer.
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

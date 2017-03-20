package deckMaker

import (
	"bufio"
	"log"
	"os"
)

func main() {
	in, err := os.Open("./testdata/in.txt") // FIXME
	defer in.Close()
	if err != nil {
		log.Fatalf("Error occured while opening a fle. error = %+v", err)
		os.Exit(1)
	}

	out, err := os.Create("./out/out.txt")
	defer out.Close()
	if err != nil {
		log.Fatalf("Error occured while creating a fle. error = %+v", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) == 0 {
			break
		}
		meaning, err := getMeaning(word)
		if err != nil {
			log.Fatalf("Error occured writing to file. error: %v", err)
		}

		if err := appendToFile(out, word, meaning); err != nil {
			log.Fatalf("Error occured writing to file. error: %v", err)
		}
		if err := scanner.Err(); err != nil {
			log.Printf("%+v", err)
		}
	}
}

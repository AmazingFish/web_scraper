package deckMaker

import "io"

func appendToFile(file io.Writer, word string, definition string) error {
	_, err := file.Write([]byte(word + "\t" + definition + "\n"))
	return err
}

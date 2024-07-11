package ch2

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

// Expand the program you wrote in the first exercise so that instead of printing the contents of the text files,
// it searches for a string match.
// The string to search for is the first argument on the command line.
// When you spawn a new goroutine, instead of printing the file’s contents, it should read the file and search for a match.
// If the goroutine finds a match, it should output a message saying that the filename contains a match.
// Call the program grepfiles.go.
// Here’s how you can execute this Go program
// (“bubbles” is the search string in this example):
// go run grepfiles.go bubbles txtfile1 txtfile2 txtfile3

func getFileContentsBytes(filename string) ([]byte, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}
	return b, nil
}

func matchString(searchTerm, filename string) (bool, error) {
	c, err := getFileContentsBytes(filename)
	if err != nil {
		return false, err
	}
	found := bytes.Contains(c, []byte(searchTerm))
	return found, nil
}

func Grepfiles() {
	searchTerm := os.Args[1:2]
	filenames := os.Args[2:]
	for i := 0; i < len(filenames); i++ {
		go func(serachTerm, filename string) {
			ok, err := matchString(serachTerm, filename)
			if err != nil {
				fmt.Printf("Error trying to match string %s on file %s: %v\n", serachTerm, filename, err)
			}
			if ok {
				fmt.Printf("Word %s found for file %s\n", serachTerm, filename)
			}

		}(searchTerm[0], filenames[i])
	}
	time.Sleep(2 * time.Second)
}

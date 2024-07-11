package ch2

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Change the program you wrote in the second exercise so that instead of passing a list of text filenames,
// you pass a directory path.
// The program will look inside this directory and list the files.
// For each file, you can spawn a goroutine that will search for a string match (the same as before).
// Call the program grepdir.go.
// Here’s how you can execute this Go program:
// go run grepdir.go bubbles ../../commonfiles

func getDirContents(dirName string) ([]string, error) {
	contents := []string{}
	c, err := os.ReadDir(dirName)
	if err != nil {
		return []string{}, err
	}
	for _, v := range c {
		contents = append(contents, v.Name())
	}
	return contents, nil
}

func Grepdir() {
	searchTerm := os.Args[1]
	dirName := os.Args[2]

	filenames, err := getDirContents(dirName)
	if err != nil {
		fmt.Printf("Failed to get dir contents with err %v\n", err)
	}
	for _, f := range filenames {
		go func(searchTerm, v string) {
			found := strings.Contains(searchTerm, f)
			if found {
				fmt.Printf("Found search term %s on directory %s\n", searchTerm, dirName)
			}
		}(searchTerm, f)
	}
	time.Sleep(2 * time.Second)
}

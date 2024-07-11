package ch2

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Adapt the program in the third exercise to continue searching recursively in any subdirectories.
// If you give your search goroutine a file, it should search for a string match in that file,
// just like in the previous exercises.
// Otherwise, if you give it a directory,
// it should recursively spawn a new goroutine for each file or directory found inside.
// Call the program grepdirrec.go, and execute it by running this command:
// go run grepdirrec.go bubbles ../../commonfiles

func isDir(filename string) (bool, error) {
	f, err := os.Open(filename)
	if err != nil {
		return false, err
	}

	i, err := f.Stat()
	if err != nil {
		return false, err
	}

	return i.IsDir(), nil
}

func searchInFile(searchTerm, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open file %s with error %v\n", filename, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), searchTerm) {
			fmt.Printf("Found search term %s in file %s\n", searchTerm, filename)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
	}
}

func processDir(searchTerm, dirName string) {
	contents, err := os.ReadDir(dirName)
	if err != nil {
		fmt.Printf("Failed to read directory %s with error %v\n", dirName, err)
		return
	}

	for _, entry := range contents {
		entryPath := filepath.Join(dirName, entry.Name())
		if entry.IsDir() {
			go processDir(searchTerm, entryPath)
			time.Sleep(2 * time.Second)
		} else {
			go searchInFile(searchTerm, entryPath)
			time.Sleep(2 * time.Second)
		}
	}
}
func Grepdirrec() {
	searchTerm := os.Args[1]
	dirName := os.Args[2]

	go processDir(searchTerm, dirName)
	time.Sleep(2 * time.Second)

}

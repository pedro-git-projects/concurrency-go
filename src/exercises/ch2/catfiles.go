package ch2

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

// Write a program similar to the one in listing 2.3 that accepts a list of text filenames as arguments.
// For each filename, the program should spawn a new goroutine that will output the contents of that file to the console.
// You can use the time.Sleep() function to wait for the child goroutines to complete (until you know how to do this better).
// Call the program catfiles.go.
// Here’s how you can execute this Go program:

// go run catfiles.go txtfile1 txtfile2 txtfile3

func getFileContentsStr(filename string) (string, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	content := bytes.NewBuffer(b).String()
	return content, nil
}

func printFileContents(filename string) error {
	c, err := getFileContentsStr(filename)
	if err != nil {
		return err
	}
	fmt.Println(c)
	return nil
}

func Catfiles() {
	filenames := os.Args[1:]
	for i := 0; i < len(filenames); i++ {
		go func(filename string) {
			if err := printFileContents(filename); err != nil {
				fmt.Printf("Failed to read file contentes with err %s\n", err)
			}
		}(filenames[i])
	}
	time.Sleep(2 * time.Second)
}

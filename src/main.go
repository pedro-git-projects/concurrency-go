package main

import (
	"log"
	"strings"
)

func main() {
	app := New()
	app.Greet()
	option, err := app.reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	app.Run(strings.TrimSpace(option))
}

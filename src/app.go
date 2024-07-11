package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pedro-git-projects/concurrency-go/src/exercises/ch2"
)

type App struct {
	exerciseList ExerciseList
	reader       *bufio.Reader
}

func New() *App {
	list := NewExerciseList()
	list.AddExercise("catfiles", ch2.Catfiles)
	list.AddExercise("grepfiles", ch2.Grepfiles)
	list.AddExercise("grepdir", ch2.Grepdir)
	list.AddExercise("grepdirrec", ch2.Grepdirrec)
	return &App{
		exerciseList: list,
		reader:       bufio.NewReader(os.Stdin),
	}
}

func (app App) Greet() {
	fmt.Printf("Welcome, please chose an exercise\n")
	for k := range app.exerciseList.funcMap {
		fmt.Printf("-> %s\n", k)
	}
}

func (app App) Run(exercise string) {
	app.exerciseList.Run(exercise)
}

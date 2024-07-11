package main

import (
	"fmt"
	"log"
)

type exerciseFunc func()

type ExerciseList struct {
	funcMap map[string]exerciseFunc
}

func NewExerciseList() ExerciseList {
	return ExerciseList{
		funcMap: make(map[string]exerciseFunc),
	}
}

func (e *ExerciseList) AddExercise(name string, fn exerciseFunc) {
	e.funcMap[name] = fn
}

func (e *ExerciseList) Run(name string) {
	fn, ok := e.funcMap[name]
	if !ok {
		log.Printf("Exercise %s not found, please enter a valid name such as: ", name)
		for k := range e.funcMap {
			fmt.Println(k)
		}
		panic("invalid exercise")
	}
	fn()
}

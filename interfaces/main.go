package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
}

type englistBot struct{}
type spanishBot struct{}

func main() {
	eb := englistBot{}
	sp := spanishBot{}

	printGreeting(eb)
	printGreeting(sp)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englistBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

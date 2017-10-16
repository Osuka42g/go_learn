package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishhBot struct{}

func main() {
	eb := englishBot{}
	sp := spanishhBot{}

	printGreeting(eb)
	printGreeting(sp)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi there"
}

func (spanishhBot) getGreeting() string {
	return "Hola!"
}

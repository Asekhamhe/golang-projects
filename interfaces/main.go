package main

import "fmt"

type bot interface {
	getGreetings() string
	// getTestInt() int
}

type englishBot struct{}

type spanishBot struct{}

type testString string

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	var tb testString

	printGreeting(eb)
	printGreeting(sb)
	printGreeting(tb)

}

func (englishBot) getGreetings() string {
	// different logic for bots ==> Custom logic
	return "Hi there"
}

func (spanishBot) getGreetings() string {
	// custom logic
	return "Hola!"
}

func (testString) getGreetings() string {
	return "Test string"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreetings())
}

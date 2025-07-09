package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	gpt := englishBot{}
	sgpt := spanishBot{}
	printGreeting(gpt)
	printGreeting(sgpt)
}
func printGreeting(eb bot) {
	fmt.Println(eb.getGreeting())
}
func (eb englishBot) getGreeting() string {
	return "Hi there"
}
func (sb spanishBot) getGreeting() string {
	return "Hola"
}

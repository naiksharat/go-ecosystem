package main

import "fmt"

type spanishbot struct{}

type englishBot struct{}

func (eb englishBot) getGreeting() string { //func (englishBot) getGreeting() string // since eb is not used, you can omit it
	return "hi there"
}

func (sb spanishbot) getGreeting() string {
	return "hola"
}

func printGreetingEnglish(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreetingSpanish(sb englishBot) {
	fmt.Println(sb.getGreeting())
}

func main() {

}

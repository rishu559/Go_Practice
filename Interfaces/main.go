package main

import "fmt"

type egBot struct{}
type spBot struct{}

type bot interface{
	getGreetings() string
}

func main(){
	eb := egBot{}
	sp := spBot{}

	printGreeting(eb)
	printGreeting(sp)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreetings())
}

func (eg egBot) getGreetings() string{
	return "hello"
}

func (eg spBot) getGreetings() string{
	return "hola!"
}


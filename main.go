package main

import "fmt"

const helloWorld = "Hello World "

func Hello(nome string) string {
	if nome == "" {
		nome = "for you"
	}
	return helloWorld + nome
}

func main() {
	fmt.Println(Hello("Lucas"))
}
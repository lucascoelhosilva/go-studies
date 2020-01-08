package main

import "fmt"

func Hello(nome string) string {
	return "Hello World " + nome
}

func main() {
	fmt.Println(Hello("Lucas"))
}
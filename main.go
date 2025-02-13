package main

import "fmt"

const test_const = "Teste, "

func Hello(name string) string {
	return test_const + name
}

func main() {
	fmt.Println(Hello("Golang!"))
}

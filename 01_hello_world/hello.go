package main 

import "fmt"

const englishHelloPrifix = "Hello, "

func Hello(name string) string{
	return englishHelloPrifix + name
}

func main() {
	fmt.Println(Hello("world"))
}
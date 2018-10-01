package main

import "fmt"

//Separate domain code from external side effects

const helloprefix = "hello, "

func Hello(name string) string {
	return helloprefix + name
}

func main() {
	fmt.Printf(Hello("Tinashes"))
}

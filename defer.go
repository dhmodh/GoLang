package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

/*
A defer statement defers the execution of a function until the surrounding function returns.
*/

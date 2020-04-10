package main

import "fmt"

func fact(n int) int {

	if n > 0 {
		var result = n * fact(n-1)
		return result
	}
	return 1
}

func main() {
	var num int
	fmt.Println("Enter the Number to Find Factorial:- ")
	fmt.Scanf("%d", &num)
	fmt.Println(fact(num))
}

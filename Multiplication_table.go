package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter the Number:- ")
	fmt.Scanf("%d %d", &num)
	for i := 1; i <= 10; i++ {
		fmt.Println(num, "*",i, "=",num*i)
	}
}

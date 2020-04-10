package main

import "fmt"

func main() {
	var size int
	fmt.Println("Enter the size of Array do u want make:- ")
	fmt.Scanf("%d", &size)
	arr := make([]int, size)
	fmt.Println("Enter the Number:- ")
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	var max = arr[0]
	for i := 1; i < size; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	fmt.Println(max, "is Max out of", arr)
}

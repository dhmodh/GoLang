package main

import (
	"fmt"
	"sort"
)

func main() {
	var size int
	fmt.Println("Enter the Size of Array:- ")
	fmt.Scanf("%d", &size)
	arr := make([]int, size)
	fmt.Println("Enter the Number of Array:- ")
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Println("Current Array:- ", arr)
	sort.Ints(arr)
	fmt.Println("After Sorting:- ", arr)
}

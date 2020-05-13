/*
Simple Intex
	func functionname(parametername type) return type{
		function body
	}
*/

package main

import "fmt"

func add(x int, y int) int {  		//or func add(x,y int) int { }
	return x + y
}

func main() {
	fmt.Println(add(42,13))
}

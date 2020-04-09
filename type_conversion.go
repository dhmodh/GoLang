package main

import(
	"fmt"
)

func main(){
	i := 55		//int
	j := 67.8		//float
	sum := i + int(j)		//J is converted to int
	fmt.Println(sum)
}

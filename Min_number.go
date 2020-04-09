package main

import (
	"fmt"
	"math"
)

func main(){
	a,b :=145.8,543.78
	c := math.Min(a,b)
	fmt.Println("Minimum value is:- ",c)	
}

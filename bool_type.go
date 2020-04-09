package main

import "fmt"

func main() {  
    a := true
    b := false
    fmt.Println("a:", a, "b:", b)
    c := a && b		//The && operator returns true only when both a and b are true. 
    fmt.Println("c:", c)
    d := a || b		//The || operator returns true when either a or b is true.
    fmt.Println("d:", d)
}

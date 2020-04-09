package main

import (  
    "fmt"
    "unsafe"	//Go has a package unsafe which has a Sizeof function which returns in bytes the size of the variable passed to it.
)

func main() {  
    var a int = 89
    b := 95
    fmt.Println("value of a is", a, "and b is", b)
    fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a)) 			//type and size of a
    fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b)) 		//type and size of b
}

//%T is the format specifier to print the type and %d is used to print the size.

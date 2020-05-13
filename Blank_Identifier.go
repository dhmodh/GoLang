package main

import "fmt"

func rectProps(lenght,width float64)(float64,float64){
	var area = lenght * width
	var perimeter = (lenght + width)*2
	return area, perimeter
}

func main(){
	area, _ := rectProps(10.8,5.6) 		//Perimeter is Discarted
	fmt.Println("Area is ",area)
}

/*
	_ is know as the blank identifier in GoLang
	It can be use to in place  of any value of any type.
*/

package main

import (
	"fmt"
)

func calculateBill(price, no int) int {
	var totalprice = price * no
	return totalprice
}

func main() {
	price, no := 90, 6
	totalprice := calculateBill(price, no)
	fmt.Println("Total price is:- ", totalprice)
}

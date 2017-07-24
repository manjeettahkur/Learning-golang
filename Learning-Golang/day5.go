package main

import (
	"fmt"
)

func calculatedBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice

}
func main() {
	total := calculatedBill(10, 2)
	fmt.Println(" Total price is", total)
}

/**
 * First Day of my Learning 
 */

package main

// Import files

import (
	"fmt"
	"math"
)
// Find minimum value function

func FindSmall(a,b float64) float64 {
	return math.Min(a,b)
}

func main() {
	a := 30.00
	b := 40.00
	var smallNumber = FindSmall(a,b)
	fmt.Println("Small number is",smallNumber)
}

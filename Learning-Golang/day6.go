package main

import "fmt"

func main() {
	area, parameter := reactProps(10.8, 5.6)
	fmt.Printf(" Area %f perimeter %f", area, parameter)
}

func reactProps(length, width float64) (float64, float64) {
	var area = length * width
	var parameter = (length + width) * 2
	return area, parameter
}

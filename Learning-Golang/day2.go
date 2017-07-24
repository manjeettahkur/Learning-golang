/**
 * Second Day of my Learning
 */
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int = 85
	b := 90
	fmt.Println(" value of a is", a, " value of b is", b)
	fmt.Printf(" type of a is %T size of a is %d", a, unsafe.Sizeof(a))
	fmt.Printf("\ntype of b is %T size of a is %d", b, unsafe.Sizeof(b))
}

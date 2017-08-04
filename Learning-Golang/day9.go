package main

import "fmt"

func foo(n int) func() (int, bool) {
	i := -1
	return func() (int, bool) {
		if i >= n {
			return 0, true
		}
		i++
		return i, false
	}
}

func main() {
	f := foo(5)
	for i, eof := f(); !eof; i, eof = f() {
		fmt.Println(i)
	}
}
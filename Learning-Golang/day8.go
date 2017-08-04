package main

import "strings"

func wordCount(s string) map[string]int {
	ss := strings.Fields(s)
	num := len(ss)
	ret := make(map[string]int)

	for i := 0; i < num; i++ {
		(ret[ss[i]])++

	}
}

func main() {
}

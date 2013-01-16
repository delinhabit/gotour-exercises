package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	x := make(map[string]int)
	for i := 0; i < len(fields); i++ {
		x[fields[i]]++
	}
	return x
}

func main() {
	wc.Test(WordCount)
}

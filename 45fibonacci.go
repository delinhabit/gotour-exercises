package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	an2 := 0
	an1 := 1
	return func() int {
		ret := an1
		an := an1 + an2
		an2 = an1
		an1 = an
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

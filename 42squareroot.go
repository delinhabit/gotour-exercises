package main

import (
        "fmt"
        "math"
)

func Sqrt(x float64) (float64, int) {
        z := float64(x / 2)
	last := float64(0)
	i := 0

	for math.Abs(last - z) > 0.0000000001 {
                last = z
                z = z - (z*z-x)/(2*z)
		i++
        }
	
        return z, i
}

func main() {
	for x := 1; x < 20; x++ {
		z, i := Sqrt(float64(x))
        	fmt.Printf("sqrt(%d) = %f (%d iterations)\n", x, z, i)
	}
}

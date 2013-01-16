package main

import "fmt"
import "math/cmplx"

func Cbrt(x complex128) (complex128, int) {
	z := complex128(x / 2)
	last := complex128(0)
	i := 0
	
	for cmplx.Abs(last - z) > 0.0000000001 {
		last = z
		z = z - (z*z*z - x) / (3*z*z)
		i++
	}
	
	return z, i
}

func main() {
	var nr complex128 = 2
	croot, it := Cbrt(nr)
	fmt.Printf("Cube root of %g = %g (%d iterations)\n", nr, croot, it)
	fmt.Printf("Validation: %g ^ 3 = %g\n", croot, cmplx.Pow(croot, 3))
}

package main

import "code.google.com/p/go-tour/pic"

func power(x, y int) uint8 {
	return uint8(x ^ y)
}

func median(x, y int) uint8 {
	return uint8((x + y) / 2)
}

func mult(x, y int) uint8 {
	return uint8(x * y)
}

func Pic(f func(int, int) uint8) func(int, int) [][]uint8 {
	return func(dx, dy int) [][]uint8 {
		pic := make([][]uint8, dy)
		for i := 0; i < dy; i++ {
			pic[i] = make([]uint8, dx)
		}

		for y := 0; y < dy; y++ {
			for x := 0; x < dx; x++ {
				pic[x][y] = f(x, y)
			}
		}
		return pic
	}
}

func main() {
	pic.Show(Pic(power))
}

package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type Image struct {
	w, h int
	f    func(int, int) uint8
}

func main() {
	m := Image{255, 255, power}
	pic.ShowImage(m)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) At(x, y int) color.Color {
	v := i.f(x, y)
	return color.RGBA{v, v, 255, 255}
}

func power(x, y int) uint8 {
	return uint8(x ^ y)
}

func median(x, y int) uint8 {
	return uint8((x + y) / 2)
}

func mult(x, y int) uint8 {
	return uint8(x * y)
}

package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	h int
	w int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

func (img Image) At(x, y int) color.Color {
	// ここの式は適当に決めてOK
	v := uint8((x + y) / 2)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{100, 100}
	pic.ShowImage(m)
}

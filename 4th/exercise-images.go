package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type MyImage struct{}

func (i MyImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (i MyImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, 200, 100)
}

func (i MyImage) At(x int, y int) color.Color {
	return color.RGBA{0, 150, 255, 255}
}

func main() {
	m := MyImage{}
	pic.ShowImage(m)
}

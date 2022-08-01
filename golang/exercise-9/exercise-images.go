package main

import (
	"golang.org/x/tour/pic"
	"image/color"
	"image"
)

type Image struct{}

func (_ Image)ColorModel()color.Model{
	return color.RGBAModel
}

func (_ Image)Bounds()image.Rectangle{
	return image.Rect(0,0,255,255)
}

func (img Image)At(x, y int) color.Color{
	return color.RGBA{uint8(x),uint8(y),255,255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}

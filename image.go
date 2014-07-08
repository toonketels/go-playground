package main

import (
    "code.google.com/p/go-tour/pic"
    "image"
    "image/color"
)

type Image struct{
	Width int
    Height int
}


func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i *Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, i.Width, i.Height)

}

// Returns the color at a given x, y coordintate...
// We just specify it as a value in func of the coordinate.
func (i *Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}


func main() {
    m := &Image{200, 200}
    pic.ShowImage(m)
}

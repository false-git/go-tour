package main

import (
    "code.google.com/p/go-tour/pic"
    "image"
    "image/color"
)

type Image struct{
    w, h int
    
}
func (i Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, i.w, i.h)
}
func (i Image) ColorModel() color.Model {
    return color.RGBAModel
}
func (i Image) At(x, y int) color.Color {
    return color.RGBA{0, uint8(x), uint8(y), 255}
}

func main() {
    m := Image{256, 256}
    pic.ShowImage(m)
}

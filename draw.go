package main

import "image/color"

func drawSquare(c color.Color, r Rectangle) {
	drawLine(c, r.Min, Point{X: r.Min.X, Y: r.Max.Y})
	drawLine(c, r.Min, Point{X: r.Max.X, Y: r.Min.Y})
	drawLine(c, Point{X: r.Min.X, Y: r.Max.Y}, r.Max)
	drawLine(c, Point{X: r.Max.X, Y: r.Min.Y}, r.Max)
}

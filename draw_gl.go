// +build !js

package main

import (
	"image/color"

	"github.com/go-gl/gl/v2.1/gl"
)

func setDisplaySize(width, height int) {
	gl.Ortho(0, float64(width), 0, float64(height), 0, 1)
}

func clearScreen() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func drawLine(c color.Color, start, end Point) {
	//gl.LineWidth(10)
	r, g, b, a := c.RGBA()
	gl.Color4us(uint16(r), uint16(g), uint16(b), uint16(a))
	gl.Begin(gl.LINES)
	gl.Vertex2d(start.X, start.Y)
	gl.Vertex2d(end.X, end.Y)
	gl.End()
}

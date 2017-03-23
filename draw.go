// +build !js

package main

import (
	"image/color"

	"github.com/go-gl/gl/v2.1/gl"
)

func drawLine(c color.Color, start, end xyz) {
	gl.LineWidth(2.5)
	r, g, b, a := c.RGBA()
	gl.Color4us(uint16(r), uint16(g), uint16(b), uint16(a))
	gl.Begin(gl.LINES)
	gl.Vertex3d(start.X, start.Y, start.Z)
	gl.Vertex3d(end.X, end.Y, end.Z)
	gl.End()
}

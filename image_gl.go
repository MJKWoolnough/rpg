package main

import "github.com/go-gl/gl/v2.1/gl"

func (i *Image) Draw() {
	gl.Enable(gl.TEXTURE_2D)
	gl.Enable(gl.BLEND)
	gl.BindTexture(gl.TEXTURE_2D, i.Image.id)
	gl.Color4f(1, 1, 1, 1)
	gl.Begin(gl.QUADS)

	gl.TexCoord2f(1, 1)
	gl.Vertex2d(i.size.Min.X, i.size.Min.Y)

	gl.TexCoord2f(0, 1)
	gl.Vertex2d(i.size.Max.X, i.size.Min.Y)

	gl.TexCoord2f(0, 0)
	gl.Vertex2d(i.size.Max.X, i.size.Max.Y)

	gl.TexCoord2f(1, 0)
	gl.Vertex2d(i.size.Min.X, i.size.Max.Y)

	gl.End()
	gl.Disable(gl.BLEND)
	gl.Disable(gl.TEXTURE_2D)
}

package main

import (
	"image"

	"github.com/go-gl/gl/v2.1/gl"
)

type Texture struct {
	Bounds image.Rectangle
	id     uint32
}

func loadTexture(img *image.RGBA) *Texture {
	var texture uint32
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(img.Rect.Size().X), int32(img.Rect.Size().Y), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(img.Pix))
	return &Texture{
		Bounds: img.Bounds(),
		id:     texture,
	}
}

func (t *Texture) Delete() {
	gl.DeleteTextures(1, &t.id)
}

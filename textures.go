package main

import (
	"image"
	"image/draw"
	"io"
)

func LoadTexture(f io.Reader) (*Texture, error) {
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)
	return loadTexture(rgba), nil
}

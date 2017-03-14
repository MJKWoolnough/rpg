//+build js

package main

import (
	"image/color"

	"github.com/MJKWoolnough/gopherjs/xdom"
	"github.com/MJKWoolnough/gopherjs/xjs"
	"github.com/gopherjs/webgl"
)

type Context struct {
	ctx *webgl.Context
}

func getContext(width, height int) (*Context, error) {
	canvas := xdom.Canvas()
	canvas.Width = width
	canvas.Height = height
	xjs.Body().AppendChild(canvas)
	ctx, err := webgl.NewContext(canvas.Object, nil)
	if err != nil {
		return nil, err
	}
	return &Context{ctx: ctx}, nil
}

func (c *Context) Close() error {
	return nil
}

func (c *Context) SetColour(cl color.Color) {
	r, g, b, a := cl.RGBA()
	c.ctx.ClearColor(float32(r)/255, float32(g)/255, float32(b)/255, float32(a)/255)
	c.ctx.Clear(0)
}

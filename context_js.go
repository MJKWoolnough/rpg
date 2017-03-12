//+build js

package main

import (
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

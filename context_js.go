//+build js

package main

import (
	"github.com/MJKWoolnough/gopherjs/xdom"
	"github.com/MJKWoolnough/gopherjs/xjs"
	"github.com/gopherjs/webgl"
	"honnef.co/go/js/dom"
)

type Context struct {
	ctx *webgl.Context
}

func getContext(width, height int) (*Context, error) {
	c := make(chan struct{})
	ctx := new(Context)
	var err error
	dom.GetWindow().AddEventListener("load", false, func(dom.Event) {
		go func() {
			canvas := xdom.Canvas()
			canvas.Width = width
			canvas.Height = height
			xjs.Body().AppendChild(canvas)
			ctx.ctx, err = webgl.NewContext(canvas.Object, nil)
			close(c)
		}()
	})
	<-c
	if err != nil {
		return nil, err
	}
	return ctx, nil
}

func (c *Context) Close() error {
	return nil
}

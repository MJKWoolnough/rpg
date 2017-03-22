//+build js

package main

import (
	"github.com/MJKWoolnough/engine/windows/webgl"
	"github.com/MJKWoolnough/gopherjs/xjs"
	"honnef.co/go/js/dom"
)

var webglI *webglengine

func main() {
	dom.GetWindow().AddEventListener("load", false, func(dom.Event) {
		if err := run(); err != nil {
			xjs.Alert(err)
		}

	})
	webglI = &webgl.Instance
}

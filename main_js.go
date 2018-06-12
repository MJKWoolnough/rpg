//+build js

package main

import (
	"vimagination.zapto.org/engine/windows/webgl"
	"vimagination.zapto.org/gopherjs/xjs"
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

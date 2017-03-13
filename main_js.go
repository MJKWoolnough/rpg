//+build js

package main

import (
	"github.com/MJKWoolnough/gopherjs/xjs"
	"honnef.co/go/js/dom"
)

func main() {
	dom.GetWindow().AddEventListener("load", false, func(dom.Event) {
		if err := run(); err != nil {
			xjs.Alert(err)
		}
	})
}
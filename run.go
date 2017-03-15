package main

import (
	"math/rand"

	"github.com/MJKWoolnough/engine"
	"github.com/go-gl/gl/v3.1/gles2"
)

func run() error {
	c := engine.Config{
		Width:  800,
		Height: 600,
		Title:  "Test",
	}
	return engine.Loop(c, loop)
}

func loop(w, h int, t float64) {
	gles2.ClearColor(rand.Float32(), rand.Float32(), rand.Float32(), 1)
	gles2.Clear(gles2.COLOR_BUFFER_BIT)
	if engine.KeyPressed(engine.KeyEscape) {
		engine.Close()
	}
}

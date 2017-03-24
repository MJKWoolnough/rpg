package main

import (
	"errors"
	"image/color"

	"github.com/MJKWoolnough/engine"
	"github.com/MJKWoolnough/limage/lcolor"
)

func run() error {
	monitors := engine.GetMonitors()
	if len(monitors) == 0 {
		return errors.New("no monitor")
	}
	modes := monitors[0].GetModes()
	if len(modes) == 0 {
		return errors.New("no modes")
	}
	c := engine.Config{
		Monitor: monitors[0],
		Mode:    modes[len(modes)-1],
		Title:   "Test",
	}
	return engine.Loop(c, loop)
}

func loop(w, h int, t float64) {
	if engine.KeyPressed(engine.KeyEscape) {
		engine.Close()
		return
	}
	drawHexGrid(lcolor.RGB{R: 255})
	//setCamera()
}

type xyz struct {
	X, Y, Z float64
}

func drawSquareGrid(c color.Color) {
	for i := float64(-1); i < 1; i += 0.1 {
		drawLine(c, xyz{i, -1, 0}, xyz{i, 1, 0})
		drawLine(c, xyz{-1, i, 0}, xyz{1, i, 0})
	}
}

func drawHexGrid(c color.Color) {
	drawLine(c, xyz{0, 0, 0}, xyz{0.1, 0, 0})
	drawLine(c, xyz{0.1, 0, 0}, xyz{0.17, 0.07, 0})
	drawLine(c, xyz{0.17, 0.07, 0}, xyz{0.1, 0.14, 0})
	drawLine(c, xyz{0.1, 0.14, 0}, xyz{0, 0.14, 0})
	drawLine(c, xyz{0, 0.14, 0}, xyz{-0.07, 0.07, 0})
	drawLine(c, xyz{-0.07, 0.07, 0}, xyz{0, 0, 0})
}

type camera struct {
	Position, Facing xyz
	Yaw, Pitch, Roll float64
}

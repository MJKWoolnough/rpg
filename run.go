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
	if err := engine.Init(engine.Config{
		Monitor: monitors[0],
		Mode:    modes[len(modes)-1],
		Title:   "Test",
	}); err != nil {
		return err
	}
	engine.Loop(loop)
	return engine.Uninit()
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
	offset := true
	for j := float64(-1); j <= 1.07; j += 0.07 {
		start := float64(-1)
		if offset {
			start += 0.17
			offset = false
		} else {
			offset = true
		}
		for i := start; i <= 1.34; i += 0.34 {
			drawLine(c, xyz{i - 0.1, j, 0}, xyz{i, j, 0})
			drawLine(c, xyz{i, j, 0}, xyz{i + 0.07, j - 0.07, 0})
			drawLine(c, xyz{i, j, 0}, xyz{i + 0.07, j + 0.07, 0})
		}
	}
}

type camera struct {
	Position, Facing xyz
	Yaw, Pitch, Roll float64
}

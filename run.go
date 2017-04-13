package main

import (
	"errors"

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
	//g.typ = 1
	engine.Loop(loop)
	return engine.Uninit()
}

var (
	g      Grid
	offset Point
)

func loop(w, h int, t float64) bool {
	if engine.KeyPressed(engine.KeyEscape) {
		engine.Close()
		return false
	}
	if engine.KeyPressed(engine.KeyLeft) {
		offset.X += 0.015625
	} else if engine.KeyPressed(engine.KeyRight) {
		offset.X -= 0.015625
	}
	if engine.KeyPressed(engine.KeyUp) {
		offset.Y -= 0.015625
	} else if engine.KeyPressed(engine.KeyDown) {
		offset.Y += 0.015625
	}
	clearScreen()
	g.Draw(lcolor.RGB{R: 255}, float64(w)/float64(h), offset, Rectangle{Min: Point{-0.5, -0.5}, Max: Point{0.5, 0.5}})
	green := lcolor.RGB{G: 255}
	drawLine(green, Point{-0.5, -0.5}, Point{0.5, -0.5})
	drawLine(green, Point{-0.5, -0.5}, Point{-0.5, 0.5})
	drawLine(green, Point{-0.5, 0.5}, Point{0.5, 0.5})
	drawLine(green, Point{0.5, -0.5}, Point{0.5, 0.5})
	return true
}

type camera struct {
	Position, Facing Point
	Yaw, Pitch, Roll float64
}

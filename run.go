package main

import (
	"errors"
	_ "image/png"
	"os"

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

	f, err := os.Open("test.png")
	if err != nil {
		return err
	}
	logo.Image, err = LoadTexture(f)
	if err != nil {
		return err
	}
	f.Close()

	logo.size.Max.X = float64(logo.Image.Bounds.Max.X)
	logo.size.Max.Y = float64(logo.Image.Bounds.Max.Y)

	g.Window(Rectangle{Max: Point{X: 100, Y: 100}})
	engine.Loop(loop)
	logo.Image.Delete()
	return engine.Uninit()
}

var (
	g             = NewSquareGrid(50)
	first         = true
	logo          Image
	width, height int
)

const moveAmount = float64(2)

func loop(w, h int, t float64) bool {
	var offset Point
	render := false
	if engine.KeyPressed(engine.KeyEscape) {
		engine.Close()
		return false
	}
	if w != width || h != height {
		width = w
		height = h
		render = true
		setDisplaySize(width, height)
		g.window = Rectangle{Min: Point{X: float64(width) / 4, Y: float64(height) / 4}, Max: Point{X: 3 * float64(width) / 4, Y: 3 * float64(height) / 4}}
	}
	if engine.KeyPressed(engine.KeyLeft) {
		offset.X = moveAmount
		render = true
	} else if engine.KeyPressed(engine.KeyRight) {
		offset.X = -moveAmount
		render = true
	}
	if engine.KeyPressed(engine.KeyUp) {
		offset.Y = -moveAmount
		render = true
	} else if engine.KeyPressed(engine.KeyDown) {
		offset.Y = moveAmount
		render = true
	}
	if !render && !first {
		return false
	}
	first = false
	clearScreen()
	g.OffsetBy(offset)
	logo.OffsetBy(offset)
	g.Draw(lcolor.RGB{R: 255})
	drawSquare(lcolor.RGB{G: 255}, g.window)
	drawLine(lcolor.RGB{B: 255}, Point{}, Point{X: 0.5, Y: 0.5})
	logo.Draw()
	return true
}

type camera struct {
	Position, Facing Point
	Yaw, Pitch, Roll float64
}

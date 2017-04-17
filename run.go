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

	logo.size.Max.X = 0.2
	logo.size.Max.Y = 0.2

	g.Window(Rectangle{Min: Point{X: -0.5, Y: -0.5}, Max: Point{X: 0.5, Y: 0.5}})
	engine.Loop(loop)
	logo.Image.Delete()
	return engine.Uninit()
}

var (
	g     = NewSquareGrid(0.1)
	first = true
	logo  Image
)

const moveAmount = float64(1) / 256

func loop(w, h int, t float64) bool {
	var offset Point
	if engine.KeyPressed(engine.KeyEscape) {
		engine.Close()
		return false
	}
	if engine.KeyPressed(engine.KeyLeft) {
		offset.X = moveAmount
	} else if engine.KeyPressed(engine.KeyRight) {
		offset.X = -moveAmount
	}
	if engine.KeyPressed(engine.KeyUp) {
		offset.Y = -moveAmount
	} else if engine.KeyPressed(engine.KeyDown) {
		offset.Y = moveAmount
	}
	if offset.Eq(Point{}) && !first {
		return false
	}
	first = false
	clearScreen()
	g.OffsetBy(offset)
	logo.OffsetBy(offset)
	g.Draw(lcolor.RGB{R: 255}, float64(w)/float64(h))
	drawSquare(lcolor.RGB{G: 255}, Rectangle{Min: Point{X: -0.5, Y: -0.5}, Max: Point{X: 0.5, Y: 0.5}})
	logo.Draw(float64(w) / float64(h))
	return true
}

type camera struct {
	Position, Facing Point
	Yaw, Pitch, Roll float64
}

package main

import "image/color"

type GridType uint8

const (
	GridSquare GridType = iota
	GridHex
	//GridIso
)

type Grid struct {
	typ GridType
}

func (g *Grid) Adjacent(x, y int) [][2]int {
	return nil
}

func (g *Grid) Draw(c color.Color) {
	switch g.typ {
	case GridHex:
		drawHexGrid(c)
	default:
		drawSquareGrid(c)
	}
}

func (g *Grid) ScreenCoords(x, y int) (int, int) {
	return 0, 0
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

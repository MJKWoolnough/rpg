package main

import (
	"image"
	"image/color"
)

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

func (g *Grid) Draw(c color.Color, ratio float64, bounds image.Rectangle) {
	var xScale, yScale float64 = 1, 1
	if ratio > 1 {
		xScale = 1 / ratio
	} else if ratio < 1 {
		yScale = ratio
	}
	switch g.typ {
	case GridHex:
		drawHexGrid(c, xScale, yScale)
	default:
		drawSquareGrid(c, xScale, yScale)
	}
}

func (g *Grid) ScreenCoords(x, y int) (int, int) {
	return 0, 0
}

func drawSquareGrid(c color.Color, xScale, yScale float64) {
	for i := float64(-1); i < 1; i += xScale / 10 {
		drawLine(c, xyz{i, -1, 0}, xyz{i, 1, 0})
	}
	for i := float64(-1); i < 1; i += yScale / 10 {
		drawLine(c, xyz{-1, i, 0}, xyz{1, i, 0})
	}
}

func drawHexGrid(c color.Color, xScale, yScale float64) {
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

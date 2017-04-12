package main

import (
	"image/color"
	"math"
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

func (g *Grid) Draw(c color.Color, ratio float64, bounds Rectangle) {
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
		drawLine(c, Point{i, -1}, Point{i, 1})
	}
	for i := float64(-1); i < 1; i += yScale / 10 {
		drawLine(c, Point{-1, i}, Point{1, i})
	}
}

const (
	hexSide float64 = 0.05
	hexX            = hexSide / 2
)

var hexY = math.Sqrt(3 * (hexSide * hexSide) / 4)

func drawHexGrid(c color.Color, xScale, yScale float64) {
	xSkip := 2 * (hexSide + hexX) * xScale
	ySkip := hexY * yScale
	offset := false
	for j := float64(-1); j <= 1+ySkip; j += ySkip {
		xStart := float64(-1)
		if offset {
			xStart -= (hexSide + hexX) * xScale
			offset = false
		} else {
			offset = true
		}
		for i := xStart; i <= 1+xSkip; i += xSkip {
			drawLine(c, Point{i - hexSide*xScale, j}, Point{i, j})
			drawLine(c, Point{i, j}, Point{i + hexX*xScale, j - hexY*yScale})
			drawLine(c, Point{i, j}, Point{i + hexX*xScale, j + hexY*yScale})
		}
	}
}

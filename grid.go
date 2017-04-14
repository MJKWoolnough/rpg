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

func (g *Grid) Draw(c color.Color, ratio float64, offset Point, window Rectangle) {
	var xScale, yScale float64 = 1, 1
	if ratio > 1 {
		xScale = 1 / ratio
	} else if ratio < 1 {
		yScale = ratio
	}
	switch g.typ {
	case GridHex:
		drawHexGrid(c, xScale, yScale, offset, window)
	default:
		drawSquareGrid(c, xScale, yScale, offset, window)
	}
}

func (g *Grid) ScreenCoords(x, y int) (int, int) {
	return 0, 0
}

const squareSide = 0.1

func drawSquareGrid(c color.Color, xScale, yScale float64, offset Point, window Rectangle) {
	startX := window.Min.X + math.Mod(offset.X, xScale*squareSide)
	if startX < window.Min.X {
		startX += xScale * squareSide
	}
	for i := startX; i <= window.Max.X; i += xScale * squareSide {
		drawLine(c, Point{i, window.Min.Y}, Point{i, window.Max.Y})
	}
	startY := window.Min.Y + math.Mod(offset.Y, yScale*squareSide)
	if startY < window.Min.Y {
		startY += yScale * squareSide
	}
	for i := startY; i <= window.Max.Y; i += yScale * squareSide {
		drawLine(c, Point{window.Min.X, i}, Point{window.Max.X, i})
	}
}

const (
	hexSide float64 = 0.05
	hexX            = hexSide / 2
)

var hexY = math.Sqrt(3 * (hexSide * hexSide) / 4)

func drawHexGrid(c color.Color, xScale, yScale float64, offset Point, window Rectangle) {
	xSkip := 2 * (hexSide + hexX) * xScale
	ySkip := hexY * yScale
	rowOffset := false
	for j := float64(-1); j <= 1+ySkip; j += ySkip {
		xStart := float64(-1)
		if rowOffset {
			xStart -= (hexSide + hexX) * xScale
			rowOffset = false
		} else {
			rowOffset = true
		}
		for i := xStart; i <= 1+xSkip; i += xSkip {
			drawLine(c, Point{i - hexSide*xScale, j}, Point{i, j})
			drawLine(c, Point{i, j}, Point{i + hexX*xScale, j - hexY*yScale})
			drawLine(c, Point{i, j}, Point{i + hexX*xScale, j + hexY*yScale})
		}
	}
}

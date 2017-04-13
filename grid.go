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

func (g *Grid) Draw(c color.Color, ratio float64, offset Point, bounds Rectangle) {
	var xScale, yScale float64 = 1, 1
	if ratio > 1 {
		xScale = 1 / ratio
	} else if ratio < 1 {
		yScale = ratio
	}
	switch g.typ {
	case GridHex:
		drawHexGrid(c, xScale, yScale, offset, bounds)
	default:
		drawSquareGrid(c, xScale, yScale, offset, bounds)
	}
}

func (g *Grid) ScreenCoords(x, y int) (int, int) {
	return 0, 0
}

const squareSide = 0.1

func drawSquareGrid(c color.Color, xScale, yScale float64, offset Point, bounds Rectangle) {
	startX := bounds.Min.X + math.Mod(offset.X, xScale*squareSide)
	if startX < bounds.Min.X {
		startX += xScale * squareSide
	}
	for i := startX; i <= bounds.Max.X; i += xScale * squareSide {
		drawLine(c, Point{i, bounds.Min.Y}, Point{i, bounds.Max.Y})
	}
	startY := bounds.Min.Y + math.Mod(offset.Y, yScale*squareSide)
	if startY < bounds.Min.Y {
		startY += yScale * squareSide
	}
	for i := startY; i <= bounds.Max.Y; i += yScale * squareSide {
		drawLine(c, Point{bounds.Min.X, i}, Point{bounds.Max.X, i})
	}
}

const (
	hexSide float64 = 0.05
	hexX            = hexSide / 2
)

var hexY = math.Sqrt(3 * (hexSide * hexSide) / 4)

func drawHexGrid(c color.Color, xScale, yScale float64, offset Point, bounds Rectangle) {
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

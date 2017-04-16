package main

import (
	"image/color"
	"math"
)

type Grid interface {
	OffsetBy(Point)
	Window(Rectangle)
	Draw(color.Color, float64)
}

type grid struct {
	offset Point
	window Rectangle
}

func (g *grid) OffsetBy(p Point) {
	g.offset = g.offset.Add(p)
}

func (g *grid) Window(r Rectangle) {
	g.window = r
}

type SquareGrid struct {
	Side float64
	grid
}

func NewSquareGrid(s float64) *SquareGrid {
	return &SquareGrid{Side: s}
}

func (s *SquareGrid) Draw(c color.Color, ratio float64) {
	xScale, yScale := scales(ratio)
	startX := s.window.Min.X + math.Mod(s.offset.X, xScale*s.Side)
	if startX < s.window.Min.X {
		startX += xScale * s.Side
	}
	for i := startX; i <= s.window.Max.X; i += xScale * s.Side {
		drawLine(c, Point{i, s.window.Min.Y}, Point{i, s.window.Max.Y})
	}
	startY := s.window.Min.Y + math.Mod(s.offset.Y, yScale*s.Side)
	if startY < s.window.Min.Y {
		startY += yScale * s.Side
	}
	for i := startY; i <= s.window.Max.Y; i += yScale * s.Side {
		drawLine(c, Point{s.window.Min.X, i}, Point{s.window.Max.X, i})
	}
}

type HexGrid struct {
	Side         float64
	xDiff, yDiff float64
	grid
}

func NewHexGrid(side float64) *HexGrid {
	return &HexGrid{
		Side:  side,
		xDiff: side / 2,
		yDiff: math.Sqrt(3 * (side * side) / 4),
	}
}

func (h *HexGrid) Draw(c color.Color, ratio float64) {
	xScale, yScale := scales(ratio)
	xSkip := 2 * (h.Side + h.xDiff) * xScale
	ySkip := h.yDiff * yScale
	rowOffset := false
	for j := float64(-1); j <= 1+ySkip; j += ySkip {
		xStart := float64(-1)
		if rowOffset {
			xStart -= (h.Side + h.xDiff) * xScale
			rowOffset = false
		} else {
			rowOffset = true
		}
		for i := xStart; i <= 1+xSkip; i += xSkip {
			drawLine(c, Point{i - h.Side*xScale, j}, Point{i, j})
			drawLine(c, Point{i, j}, Point{i + h.xDiff*xScale, j - h.yDiff*yScale})
			drawLine(c, Point{i, j}, Point{i + h.xDiff*xScale, j + h.yDiff*yScale})
		}
	}
}

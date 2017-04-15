package main

func scales(ratio float64) (xScale float64, yScale float64) {
	if ratio > 1 {
		xScale = 1 / ratio
	} else if ratio < 1 {
		yScale = ratio
	}
	return xScale, yScale
}

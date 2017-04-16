package main

func scales(ratio float64) (float64, float64) {
	var xScale, yScale float64 = 1, 1
	if ratio > 1 {
		xScale = 1 / ratio
	} else if ratio < 1 {
		yScale = ratio
	}
	return xScale, yScale
}

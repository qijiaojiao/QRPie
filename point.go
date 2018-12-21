package qrpie

import "math"

type point struct {
	x float64
	y float64
}

func newPoint(x, y int) point {
	return point{float64(x), float64(y)}
}

func pointAdd(p1 point, p2 point) point {
	return point{p1.x + p2.x, p1.y + p2.y}
}

func pointMinus(p1 point, p2 point) point {
	return point{p1.x - p2.x, p1.y - p2.y}
}

func distant(p1 point, p2 point) float64 {
	d := math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
	return d
}


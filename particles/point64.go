package particles

type Point64 struct {
	X float64
	Y float64
}

func (point *Point64) Move(vector Vector64) {
	point.X += vector.X
	point.Y += vector.Y
}

func SeparationSquared(x Point64, y Point64) float64 {
	xSeparation := x.X - y.X
	ySeparation := x.Y - y.Y

	return (xSeparation * xSeparation) + (ySeparation * ySeparation)
}

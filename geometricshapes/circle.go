package geometricshapes

import (
	"math"
)

type Circle struct {
	radius float64
}

func CreateCircle(radius float64) Circle {
	c := Circle{
		radius: radius,
	}

	return c
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

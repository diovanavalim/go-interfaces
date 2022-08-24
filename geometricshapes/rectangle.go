package geometricshapes

type Rectangle struct {
	heigth float64
	width  float64
}

func CreateRectangle(heigth, width float64) Rectangle {
	r := Rectangle{heigth, width}

	return r
}

func (r Rectangle) Area() float64 {
	return r.heigth * r.width
}

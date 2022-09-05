package geometricshapes

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {

	t.Run("Circle Area", func(t *testing.T) {
		circle := Circle{10}

		expectedArea := float64(math.Pi * circle.radius * circle.radius)

		resultedArea := circle.Area()

		compareResults(expectedArea, resultedArea, t)
	})

	t.Run("Rectangle Area", func(t *testing.T) {
		rectangle := Rectangle{10, 12}

		expectedArea := float64(rectangle.heigth * rectangle.width)

		resultedArea := rectangle.Area()

		compareResults(expectedArea, resultedArea, t)
	})
}

func compareResults(expected, received float64, t *testing.T) {
	if expected != received {
		t.Errorf("A área recebida %f não é equivalente a área esperada %f", received, expected)
	}

	t.Log("Test successfully executed!")
}

package main

import (
	"fmt"
	"interfaces/geometricshapes"
	"interfaces/interfaces"
)

func calculaArea(f interfaces.Forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.Area())
}

func main() {
	circle := geometricshapes.CreateCircle(2)
	rectangle := geometricshapes.CreateRectangle(6, 4)

	calculaArea(circle)
	calculaArea(rectangle)
}

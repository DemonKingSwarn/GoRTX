package color

import (
	"fmt"
	"rt/vec3"
)

func WriteColor(pixel_color vec3.Vec3) {
	r := pixel_color.X()
	g := pixel_color.Y()
	b := pixel_color.Z()

	var rbyte int = int(255.999 * r)
	var gbyte int = int(255.999 * g)
	var bbyte int = int(255.999 * b)

	fmt.Print(rbyte, " ", gbyte, " ", bbyte, "\n")
}

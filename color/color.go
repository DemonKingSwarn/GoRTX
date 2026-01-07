package color

import (
	"fmt"
	"rt/vec3"
	"rt/interval"
)

func WriteColor(pixel_color vec3.Vec3) {
	r := pixel_color.X()
	g := pixel_color.Y()
	b := pixel_color.Z()

	intensity := interval.New(0.000, 0.999)
	var rbyte int = int(256 * intensity.Clamp(r))
	var gbyte int = int(256 * intensity.Clamp(g))
	var bbyte int = int(256 * intensity.Clamp(b))

	fmt.Print(rbyte, " ", gbyte, " ", bbyte, "\n")
}

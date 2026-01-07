package world

import (
	"rt/camera"
	"rt/hittable"
	"rt/sphere"
	"rt/vec3"
)

func Main() {
	world := new(hittable.HittableList)

	world.Add(sphere.Sphere{Center: vec3.NewXYZ(0, 0, -1), Radius: 0.5})
	world.Add(sphere.Sphere{Center: vec3.NewXYZ(0, -100.5, -1), Radius: 100})

	camera.AspectRatio = 16.0 / 9.0
	camera.ImageWidth = 400
	camera.SamplesPerPixel = 100

	camera.Render(world)
}

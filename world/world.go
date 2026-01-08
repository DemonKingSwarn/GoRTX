package world

import (
	"rt/camera"
	"rt/hittable"
	"rt/sphere"
	"rt/material"
	"rt/vec3"
)

func Main() {
	world := new(hittable.HittableList)

	material_ground := material.Lambertian{Albedo: vec3.NewXYZ(0.8, 0.8, 0.0)}
	material_center := material.Lambertian{Albedo: vec3.NewXYZ(0.1, 0.2, 0.5)}
	material_left := material.Metal{Albedo: vec3.NewXYZ(0.8, 0.8, 0.8), Fuzz: 0.3}
	material_right := material.Metal{Albedo: vec3.NewXYZ(0.8, 0.6, 0.2), Fuzz: 1.0}

	world.Add(sphere.Sphere{Center: vec3.NewXYZ(0, -100.5, -1.0), Radius: 100, Mat: &material_ground})
	world.Add(sphere.Sphere{Center: vec3.NewXYZ(0, 0, -1.2), Radius: 0.5, Mat: &material_center})
	world.Add(sphere.Sphere{Center: vec3.NewXYZ(-1.0, 0, -1.0), Radius: 0.5, Mat: &material_left})
	world.Add(sphere.Sphere{Center: vec3.NewXYZ(1.0, 0, -1.0), Radius: 0.5, Mat: &material_right})

	camera.AspectRatio = 16.0 / 9.0
	camera.ImageWidth = 400
	camera.SamplesPerPixel = 100
	camera.MaxDepth = 50

	camera.Render(world)
}

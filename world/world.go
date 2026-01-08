package world

import (
	"rt/camera"
	"rt/constants"
	"rt/hittable"
	"rt/material"
	"rt/sphere"
	"rt/vec3"
)

func Main() {
	world := new(hittable.HittableList)

	ground_material := material.Lambertian{Albedo: vec3.NewXYZ(0.5, 0.5, 0.5)}
	world.Add(sphere.Sphere{Center: vec3.NewXYZ(0.0, -1000, 0), Radius: 1000, Mat: &ground_material})

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			choose_mat := constants.RandDouble()
			center := vec3.NewXYZ(float64(a) + 0.9 * constants.RandDouble(), 0.2, float64(b) + 0.9 * constants.RandDouble())

			if vec3.Sub(center, vec3.NewXYZ(4, 0.2, 0)).Length() > 0.9 {
				var sphere_material material.Material

				if choose_mat < 0.8 {
					// diffuse
					albedo := vec3.Mul(vec3.Random(), vec3.Random())
					sphere_material = &material.Lambertian{Albedo: albedo}
					world.Add(sphere.Sphere{Center: center, Radius: 0.2, Mat: sphere_material})
				} else if choose_mat < 0.95 {
					// metal
					albedo := vec3.RandRange(0.5, 1)
					fuzz := constants.RandDoubleRange(0, 0.5)
					sphere_material = &material.Metal{Albedo: albedo, Fuzz: fuzz}
					world.Add(sphere.Sphere{Center: center, Radius: 0.2, Mat: sphere_material}) 
				} else {
					// glass
					sphere_material = &material.Dielectric{RefractionIndex: 1.5}
					world.Add(sphere.Sphere{Center: center, Radius: 0.2, Mat: sphere_material})
				}
			}
		}
	}

	material1 := material.Dielectric{RefractionIndex: 1.5}
	world.Add(sphere.Sphere{Center: vec3.NewXYZ(0, 1, 0), Radius: 1.0, Mat: &material1})
	
	material2 := material.Lambertian{Albedo: vec3.NewXYZ(0.4, 0.2, 0.1)}
	world.Add(sphere.Sphere{Center: vec3.NewXYZ(-4, 1, 0), Radius: 1.0, Mat: &material2})
	
	material3 := material.Metal{Albedo: vec3.NewXYZ(0.7, 0.6, 0.5), Fuzz: 0.0}
	world.Add(sphere.Sphere{Center: vec3.NewXYZ(4, 1, 0), Radius: 1.0, Mat: &material3})

	camera.AspectRatio = 16.0 / 9.0
	camera.ImageWidth = 1200
	camera.SamplesPerPixel = 500
	camera.MaxDepth = 50

	camera.VFOV = 20
	camera.DefocusAngle = 0.6
	camera.FocusDist = 10.0

	camera.LookFrom = vec3.NewXYZ(13, 2, 3)
	camera.LookAt = vec3.NewXYZ(0, 0, 0)
	camera.Vup = vec3.NewXYZ(0, 1, 0)

	camera.Render(world)
}

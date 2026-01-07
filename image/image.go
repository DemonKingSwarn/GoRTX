package image

import (
	"fmt"
	"math"

	"rt/color"
	"rt/hittable"
	"rt/ray"
	"rt/sphere"
	"rt/vec3"
)

func ray_color(r ray.Ray, world hittable.Hittable) vec3.Vec3 {
	rec := new(hittable.HitRecord)
	if world.Hit(r, 0, math.Inf(1), rec) {
		return vec3.MulScalar(vec3.Add(rec.Normal, vec3.NewXYZ(1, 1, 1)), 0.5)
	}

	var unit_direction = vec3.UnitVector(r.Direction())
	var a = 0.5 * (unit_direction.Y() + 1.0)
	
	return vec3.Add(
		vec3.MulScalar(vec3.NewXYZ(1.0, 1.0, 1.0), 1.0 - a), 
		vec3.MulScalar(vec3.NewXYZ(0.5, 0.7, 1.0), a),
	)
}

func Render() {
	// image
	var aspect_ratio float64 = 16.0 / 9.0
	var image_width int = 400

	// calculate the image height, and ensure that its at least 1
	var image_height int = int(float64(image_width) / aspect_ratio)
	if image_height < 1 {
		image_height = 1
	}

	// world
	world := new(hittable.HittableList)

	world.Add(sphere.Sphere{Center: vec3.NewXYZ(0, 0, -1), Radius: 0.5})
	world.Add(sphere.Sphere{Center: vec3.NewXYZ(0, -100.5, -1), Radius: 100})

	// camera
	var focal_length float64 = 1.0
	var viewport_height float64 = 2.0
	var viewport_width float64 = viewport_height * (float64(image_width) / float64(image_height))
	var camera_center = vec3.NewXYZ(0, 0, 0)

	// calculate the vector across the horizontal and down the vertical viewport edges
	var viewport_u = vec3.NewXYZ(viewport_width, 0, 0)
	var viewport_v = vec3.NewXYZ(0, -viewport_height, 0)

	// calculate the horizontal and vertical delta vector from pixel to pixel
	var pixel_delta_u = vec3.MulScalar(viewport_u, 1/float64(image_width))
	var pixel_delta_v = vec3.MulScalar(viewport_v, 1/float64(image_height))

	// calculate the location of the upper left pixel
	var viewport_upper_left = vec3.Sub(
		vec3.Sub(
			vec3.Sub(
				camera_center,
				vec3.NewXYZ(0, 0, focal_length),
			),
			vec3.MulScalar(viewport_u, 0.5),
		),
		vec3.MulScalar(viewport_v, 0.5),
	)

	var pixel00_loc = vec3.Add(viewport_upper_left, vec3.MulScalar(vec3.Add(pixel_delta_u, pixel_delta_v), 0.5))

	// render

	fmt.Print("P3\n", image_width, " ", image_height, "\n255\n")

	for j := 0; j < image_height; j++ {
		for i := 0; i < image_width; i++ {
			var pixel_center = vec3.Add(
				pixel00_loc,
				vec3.Add(
					vec3.MulScalar(pixel_delta_u, float64(i)),
					vec3.MulScalar(pixel_delta_v, float64(j)),
				),
			)
			
			var ray_direction = vec3.UnitVector(vec3.Sub(pixel_center, camera_center))
			var r = ray.New(camera_center, ray_direction)

			var pixel_color = ray_color(r, world)

			color.WriteColor(pixel_color)
		}
	}
}

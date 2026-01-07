package camera

import (
	"fmt"

	"rt/color"
	"rt/constants"
	"rt/hittable"
	"rt/interval"
	"rt/ray"
	"rt/vec3"
)

var AspectRatio float64
var ImageWidth int
var SamplesPerPixel int

func ray_color(r ray.Ray, world hittable.Hittable) vec3.Vec3 {
	rec := new(hittable.HitRecord)
	if world.Hit(r, interval.New(0, constants.Infinity), rec) {
		return vec3.MulScalar(vec3.Add(rec.Normal, vec3.NewXYZ(1, 1, 1)), 0.5)
	}

	var unit_direction = vec3.UnitVector(r.Direction())
	var a = 0.5 * (unit_direction.Y() + 1.0)
	
	return vec3.Add(
		vec3.MulScalar(vec3.NewXYZ(1.0, 1.0, 1.0), 1.0 - a), 
		vec3.MulScalar(vec3.NewXYZ(0.5, 0.7, 1.0), a),
	)
}

func Render(world hittable.Hittable) {
	// image
	aspect_ratio := AspectRatio
	image_width := ImageWidth

	// calculate the image height, and ensure that its at least 1
	var image_height int = int(float64(image_width) / aspect_ratio)
	if image_height < 1 {
		image_height = 1
	}

	var pixel_samples_scale float64 = 1.0 / float64(SamplesPerPixel)
	
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
			var pixel_color = vec3.NewXYZ(0, 0, 0)
			for sample := 0; sample < SamplesPerPixel; sample++ {
				var r = GetRay(i, j, pixel00_loc, pixel_delta_u, pixel_delta_v, camera_center)
				pixel_color = vec3.Add(pixel_color, ray_color(r, world))
			}
			pixel_color = vec3.MulScalar(pixel_color, pixel_samples_scale)
			color.WriteColor(pixel_color)
		}
	}
}

func GetRay(i, j int, pixel00_loc, pixel_delta_u, pixel_delta_v, center vec3.Vec3) ray.Ray {
	var offset = sample_square()
	var pixel_sample = vec3.Add(pixel00_loc, vec3.Add(vec3.MulScalar(pixel_delta_u, (float64(i) + offset.X())), vec3.MulScalar(pixel_delta_v, (float64(j) + offset.Y()))))

	var ray_origin = center
	var ray_direction = vec3.Sub(pixel_sample, ray_origin)

	return ray.New(ray_origin, ray_direction)
}

func sample_square() vec3.Vec3 {
	return vec3.NewXYZ(constants.RandDouble() - 0.5, constants.RandDouble() - 0.5, 0)
}

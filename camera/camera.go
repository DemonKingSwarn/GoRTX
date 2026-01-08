package camera

import (
	"fmt"
	"math"

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
var MaxDepth int

var VFOV int
var LookFrom vec3.Point3
var LookAt vec3.Point3
var Vup vec3.Vec3

var DefocusAngle float64
var FocusDist float64

var defocus_disk_u vec3.Vec3
var defocus_disk_v vec3.Vec3

func ray_color(r ray.Ray, world hittable.Hittable, depth int) vec3.Vec3 {
	if depth <= 0 {
		return vec3.New()
	}

	rec := new(hittable.HitRecord)

	if world.Hit(r, interval.New(0.001, constants.Infinity), rec) {
		scattered := new(ray.Ray)
		attenuation := new(vec3.Vec3)
		if rec.Mat.Scatter(r, rec.ToMaterialRecord(), attenuation, scattered) {
			return vec3.Mul(*attenuation, ray_color(*scattered, world, depth - 1))
		}
		return vec3.New()
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
	//var focal_length float64 = vec3.Sub(LookFrom, LookAt).Length()
	var theta float64 = constants.DegToRad(float64(VFOV))
	var h float64 = math.Tan(theta/2)
	var viewport_height float64 = 2 * h * FocusDist
	var viewport_width float64 = viewport_height * (float64(image_width) / float64(image_height))
	var camera_center = LookFrom

	var u, v, w vec3.Vec3

	// calculate the u,v,w unit basis vectors for the camera coordinate frame
	w = vec3.UnitVector(vec3.Sub(LookFrom, LookAt))
	u = vec3.UnitVector(vec3.Cross(Vup, w))
	v = vec3.Cross(w, u)

	// calculate the vector across the horizontal and down the vertical viewport edges
	var viewport_u = vec3.MulScalar(u, viewport_width)
	var viewport_v = vec3.MulScalar(v.Neg(), viewport_height)

	// calculate the horizontal and vertical delta vector from pixel to pixel
	var pixel_delta_u = vec3.MulScalar(viewport_u, 1/float64(image_width))
	var pixel_delta_v = vec3.MulScalar(viewport_v, 1/float64(image_height))

	// calculate the location of the upper left pixel
	viewport_upper_left := vec3.Sub(
		vec3.Sub(
			vec3.Sub(camera_center, vec3.MulScalar(w, FocusDist)),
			vec3.DivScalar(viewport_u, 2),
		),
		vec3.DivScalar(viewport_v, 2),
	)
	
	var pixel00_loc = vec3.Add(viewport_upper_left, vec3.MulScalar(vec3.Add(pixel_delta_u, pixel_delta_v), 0.5))

	defocus_radius := float64(FocusDist) * math.Tan(constants.DegToRad(DefocusAngle / float64(2)))
	defocus_disk_u = vec3.MulScalar(u, defocus_radius)
	defocus_disk_v = vec3.MulScalar(v,defocus_radius)

	// render

	fmt.Print("P3\n", image_width, " ", image_height, "\n255\n")

	for j := 0; j < image_height; j++ {
		for i := 0; i < image_width; i++ {
			var pixel_color = vec3.NewXYZ(0, 0, 0)
			for sample := 0; sample < SamplesPerPixel; sample++ {
				var r = GetRay(i, j, pixel00_loc, pixel_delta_u, pixel_delta_v, camera_center)
				pixel_color = vec3.Add(pixel_color, ray_color(r, world, MaxDepth))
			}
			pixel_color = vec3.MulScalar(pixel_color, pixel_samples_scale)
			color.WriteColor(pixel_color)
		}
	}
}

func GetRay(i, j int, pixel00_loc, pixel_delta_u, pixel_delta_v, center vec3.Vec3) ray.Ray {
	var offset = sample_square()
	var pixel_sample = vec3.Add(pixel00_loc, vec3.Add(vec3.MulScalar(pixel_delta_u, (float64(i) + offset.X())), vec3.MulScalar(pixel_delta_v, (float64(j) + offset.Y()))))

	var ray_origin vec3.Vec3
	if DefocusAngle <= 0 {
		ray_origin = center
	} else {
		ray_origin = DefocusDiskSample(center)
	}
	var ray_direction = vec3.Sub(pixel_sample, ray_origin)

	return ray.New(ray_origin, ray_direction)
}

func sample_square() vec3.Vec3 {
	return vec3.NewXYZ(constants.RandDouble() - 0.5, constants.RandDouble() - 0.5, 0)
}

func DefocusDiskSample(center vec3.Vec3) vec3.Point3 {
	p := vec3.RandomInUnitDisk()
	return vec3.Add(
		center, vec3.Add(
			vec3.MulScalar(defocus_disk_u, p.E[0]),
			vec3.MulScalar(defocus_disk_v, p.E[1]),
		),
	)
}

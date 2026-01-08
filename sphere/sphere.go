package sphere

import (
	"math"

	"rt/hittable"
	"rt/interval"
	"rt/material"
	"rt/ray"
	"rt/vec3"
)

type Sphere struct {
	Center vec3.Point3
	Radius float64
	Mat material.Material
}

func (s Sphere) Hit(r ray.Ray, ray_t interval.Interval, rec *hittable.HitRecord) bool {
	oc := vec3.Sub(s.Center, r.Origin())
	
	a := r.Direction().LengthSquared()
	h := vec3.Dot(r.Direction(), oc)
	c := oc.LengthSquared() - s.Radius * s.Radius
	discriminant := h * h - a * c
	
	if discriminant < 0 {
		return false
	}

	sqrtd := math.Sqrt(discriminant)

	root := (h - sqrtd) / a

	if !ray_t.Surrounds(root) {
		root = (h + sqrtd) / a
		if !ray_t.Surrounds(root) {
			return false
		}
	}

	rec.T = root
	rec.P = r.At(rec.T)
	outward_normal := vec3.DivScalar(vec3.Sub(rec.P, s.Center), s.Radius)
	rec.SetFaceNormal(r, outward_normal)
	rec.Mat = s.Mat

	return true
}

func New(center vec3.Point3, radius float64, mat material.Material) *Sphere {
	if radius < 0 {
		radius = 0
	}

	return &Sphere{
		Center: center,
		Radius: radius,
		Mat:    mat,
	}
}

package sphere

import (
	"math"

	"rt/hittable"
	"rt/ray"
	"rt/vec3"
)

type Sphere struct {
	Center vec3.Point3
	Radius float64
}

func (s Sphere) Hit(r ray.Ray, tMin float64, tMax float64, rec *hittable.HitRecord) bool {
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

	if root <= tMin || tMax <= root {
		root = (h + sqrtd) / a
		if root <= tMin || tMax <= root {
			return false
		}
	}

	rec.T = root
	rec.P = r.At(rec.T)
	outward_normal := vec3.DivScalar(vec3.Sub(rec.P, s.Center), s.Radius)
	//rec.Normal = outward_normal
	rec.SetFaceNormal(r, outward_normal)

	return true
}

package ray

import (
	"rt/vec3"
)

type Ray struct {
	Orig vec3.Point3
	Dir vec3.Vec3
}

func New(origin vec3.Point3, direction vec3.Vec3) Ray {
	return Ray{
		Orig: origin,
		Dir: direction,
	}
}

func (r Ray) Origin() vec3.Point3 {
	return r.Orig
}

func (r Ray) Direction() vec3.Vec3 {
	return r.Dir
}

func (r Ray) At(t float64) vec3.Point3 {
	return vec3.Add(r.Orig, vec3.MulScalar(r.Dir, t))
}

package material

import (
	"rt/ray"
	"rt/vec3"
)

type HitRecordLite struct {
	P         vec3.Point3
	Normal    vec3.Vec3
	FrontFace bool
	T         float64
}

type Material interface {
	Scatter(rIn ray.Ray, rec HitRecordLite, attenuation *vec3.Vec3, scattered *ray.Ray) bool
}

type Lambertian struct {
	Albedo vec3.Vec3
}

type Metal struct {
	Albedo vec3.Vec3
	Fuzz float64
}

func NewMetal(albedo vec3.Vec3) *Metal {
	return &Metal{Albedo: albedo}
}

func NewLambertian(albedo vec3.Vec3) *Lambertian {
	return &Lambertian{Albedo: albedo}
}

func (l *Lambertian) Scatter(
	rIn ray.Ray,
	rec HitRecordLite,
	attenuation *vec3.Vec3,
	scattered *ray.Ray,
) bool {

	scatter_direction := vec3.Add(rec.Normal, vec3.RandUnitVector())

	if scatter_direction.NearZero() {
		scatter_direction = rec.Normal
	}

	*scattered = ray.New(rec.P, scatter_direction)
	*attenuation = l.Albedo

	return true
}

func (m *Metal) Scatter(
	rIn ray.Ray,
	rec HitRecordLite,
	attenuation *vec3.Vec3,
	scattered *ray.Ray,
) bool {
		
	if m.Fuzz < 1 {
		m.Fuzz = 1
	}

	reflected := vec3.Reflect(rIn.Direction(), rec.Normal)
	reflected = vec3.Add(vec3.UnitVector(reflected), vec3.MulScalar(vec3.RandUnitVector(), m.Fuzz))

	*scattered = ray.New(rec.P, reflected)
	*attenuation = m.Albedo

	return (vec3.Dot(scattered.Direction(), rec.Normal) > 0)
}

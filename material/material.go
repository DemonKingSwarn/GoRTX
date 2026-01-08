package material

import (
	"math"
	"rt/constants"
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

type Dielectric struct {
	RefractionIndex float64
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

func (d *Dielectric) Scatter(
	rIn ray.Ray,
	rec HitRecordLite,
	attenuation *vec3.Vec3,
	scattered *ray.Ray,
) bool {

	*attenuation = vec3.NewXYZ(1.0, 1.0, 1.0)
	
	var refractionRatio float64
	if rec.FrontFace {
		refractionRatio = 1.0 / d.RefractionIndex
	} else {
		refractionRatio = d.RefractionIndex
	}

	unit_direction := vec3.UnitVector(rIn.Direction())
	cos_theta := math.Min(vec3.Dot(unit_direction.Neg(), rec.Normal), 1.0)
	sin_theta := math.Sqrt(1.0 - cos_theta * cos_theta)

	cannot_refract := refractionRatio * sin_theta > 1.0
	var direction vec3.Vec3

	if cannot_refract || reflectance(cos_theta, refractionRatio) > constants.RandDouble(){
		direction = vec3.Reflect(unit_direction, rec.Normal)
	} else {
		direction = vec3.Refract(unit_direction, rec.Normal, refractionRatio)
	}

	*scattered = ray.New(rec.P, direction)
	return true
}

func reflectance(cosine, refraction_index float64) float64 {
	r0 := (1 - refraction_index) / (1 + refraction_index)
	r0 = r0*r0
	return r0 + (1-r0) * math.Pow((1 - cosine), 5)
}

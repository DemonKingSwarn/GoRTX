package vec3

import (
	"math"
	"rt/constants"
)

type Vec3 struct {
	E [3]float64
}

type Point3 = Vec3

func New() Vec3 {
	return Vec3{E: [3]float64{0, 0, 0}}
}

func NewXYZ(x, y, z float64) Vec3 {
	return Vec3{E: [3]float64{x, y, z}}
}

func (v Vec3) X() float64 { return v.E[0] }
func (v Vec3) Y() float64 { return v.E[1] }
func (v Vec3) Z() float64 { return v.E[2] }

func (v Vec3) Neg() Vec3 {
	return Vec3{E: [3]float64{-v.E[0], -v.E[1], -v.E[2]}}
}

func (v Vec3) At(i int) float64 {
	return v.E[i]
}

func (v *Vec3) Set(i int, value float64) {
	v.E[i] = value
}

func (v *Vec3) AddSign(u Vec3) {
	v.E[0] += u.E[0]
	v.E[1] += u.E[1]
	v.E[2] += u.E[2]
}

func (v *Vec3) Mul(t float64) {
	v.E[0] *= t
	v.E[1] *= t
	v.E[2] *= t
}

func Add(u, v Vec3) Vec3 {
	return Vec3{E: [3]float64{
		u.E[0] + v.E[0],
		u.E[1] + v.E[1],
		u.E[2] + v.E[2],
	}}
}

func Sub(u, v Vec3) Vec3 {
	return Vec3{E: [3]float64{
		u.E[0] - v.E[0],
		u.E[1] - v.E[1],
		u.E[2] - v.E[2],
	}}
}

func Mul(u, v Vec3) Vec3 {
	return Vec3{E: [3]float64{
		u.E[0] * v.E[0],
		u.E[1] * v.E[1],
		u.E[2] * v.E[2],
	}}
}

func Dot(u, v Vec3) float64 {
	return u.E[0]*v.E[0] +
	       u.E[1]*v.E[1] +
	       u.E[2]*v.E[2]
}

func Cross(u, v Vec3) Vec3 {
	return Vec3{E: [3]float64{
		u.E[1]*v.E[2] - u.E[2]*v.E[1],
		u.E[2]*v.E[0] - u.E[0]*v.E[2],
		u.E[0]*v.E[1] - u.E[1]*v.E[0],
	}}
}

func AddScalar(v Vec3, t float64) Vec3 {
	return Vec3{E: [3]float64{
		v.E[0] + t,
		v.E[1] + t,
		v.E[2] + t,
	}}
}

func SubScalar(v Vec3, t float64) Vec3 {
	return Vec3{E: [3]float64{
		v.E[0] - t,
		v.E[1] - t,
		v.E[2] - t,
	}}
}

func MulScalar(v Vec3, t float64) Vec3 {
	return Vec3{E: [3]float64{
		v.E[0] * t,
		v.E[1] * t,
		v.E[2] * t,
	}}
}

func DivScalar(v Vec3, t float64) Vec3 {
	return MulScalar(v, 1/t)
}

func (v *Vec3) DivAssign(t float64) {
	v.Mul(1 / t)
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v Vec3) LengthSquared() float64 {
	return v.E[0]*v.E[0] + v.E[1]*v.E[1] + v.E[2]*v.E[2]
}

func UnitVector(v Vec3) Vec3 {
    return DivScalar(v, v.Length());
}

func RandUnitVector() Vec3 {
	for {
		var p = RandRange(-1, 1)
		var lensq = p.LengthSquared()
		if lensq > 1e-160 && lensq <= 1 {
			return DivScalar(p, math.Sqrt(lensq))
		}
	}
}

func RandOnHemisphere(normal Vec3) Vec3 {
	var on_unit_sphere = RandUnitVector()
	if Dot(on_unit_sphere, normal) > 0.0 {
		return on_unit_sphere
	} else {
		return on_unit_sphere.Neg()
	}
}

func Random() Vec3 {
	return NewXYZ(constants.RandDouble(), constants.RandDouble(), constants.RandDouble())
}

func RandRange(min, max float64) Vec3 {
	return NewXYZ(constants.RandDoubleRange(min, max), constants.RandDoubleRange(min, max), constants.RandDoubleRange(min, max))
}

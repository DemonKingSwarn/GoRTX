package hittable

import (
	"rt/ray"
	"rt/vec3"
)

type HitRecord struct {
	P vec3.Point3
	Normal vec3.Vec3
	T float64
	FrontFace bool
}

type Hittable interface {
	Hit(r ray.Ray, tMin, tMax float64, rec *HitRecord) bool

}

type HittableList struct {
	Objects []Hittable
}

func (rec *HitRecord) SetFaceNormal(r ray.Ray, outwardNormal vec3.Vec3) {
	rec.FrontFace = vec3.Dot(r.Direction(), outwardNormal) < 0
	if rec.FrontFace {
		rec.Normal = outwardNormal
	} else {
		rec.Normal = outwardNormal.Neg()
	}
}

func NewHittableList() *HittableList {
	return &HittableList{}
}

func NewHittableListWith(obj Hittable) *HittableList {
	list := &HittableList{}
	list.Add(obj)
	return list
}

func (hl *HittableList) Clear() {
	hl.Objects = nil
}

func (hl *HittableList) Add(obj Hittable) {
	hl.Objects = append(hl.Objects, obj)
}

func (hl *HittableList) Hit(
	r ray.Ray,
	tMin float64,
	tMax float64,
	rec *HitRecord,
) bool {

	tempRec := HitRecord{}
	hitAnything := false
	closestSoFar := tMax

	for _, object := range hl.Objects {
		if object.Hit(r, tMin, closestSoFar, &tempRec) {
			hitAnything = true
			closestSoFar = tempRec.T
			*rec = tempRec
		}
	}

	return hitAnything
}

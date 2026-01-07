package interval

import "rt/constants"

type Interval struct {
	Min float64
	Max float64
}


func NewEmpty() Interval {
	return Interval{
		Min: constants.Infinity,
		Max: -constants.Infinity,
	}
}

func New(min, max float64) Interval {
	return Interval{
		Min: min,
		Max: max,
	}
}

func (i Interval) Size() float64 {
	return i.Max - i.Min
}

func (i Interval) Contains(x float64) bool {
	return i.Min <= x && x <= i.Max
}

func (i Interval) Surrounds(x float64) bool {
	return i.Min < x && x < i.Max
}

func (i Interval) Clamp(x float64) float64 {
	if x < i.Min { return i.Min }
	if x > i.Max { return i.Max }
	return x
}

package constants

import (
	"math"
	"math/rand/v2"
)

var Infinity float64 = math.Inf(1)
var Pi float64 = 3.1415926535897932385

func DegToRad(degrees float64) float64 {
	return degrees * Pi / 180.0
}

func RandDouble() float64 {
	return rand.Float64()
}

func RandDoubleRange(min, max float64) float64 {
	return min + (max-min)*RandDouble()
}

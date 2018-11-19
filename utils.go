package sdkGO

import "math/rand"

func Random(min float64, max float64) float64 {
	return min + rand.Float64() * (max - min)
}
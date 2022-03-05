package utils

import (
	"math"
	"math/rand"
)

func RandomBoolean(chance float32) bool {
	return rand.Float32() < chance
}

func RandomItemFromInt(list []int) int {
	return list[rand.Intn(len(list))]
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

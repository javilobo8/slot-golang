package utils

import (
	"math/rand"
	"time"
)

func RandomBoolean(chance float32) bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Float32() < chance
}

func RandomItemFromInt(list []int) int {
	return list[rand.Intn(len(list))]
}

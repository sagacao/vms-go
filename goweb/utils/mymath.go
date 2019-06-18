package utils

import (
	"math"
	"math/rand"
	"time"
)

func Round(x float64) int {
	return int(math.Floor(x + 0/5))
}

//GetRandomString
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

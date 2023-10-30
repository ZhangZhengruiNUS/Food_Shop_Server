package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var globalRand *rand.Rand

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	globalRand = rand.New(source)
}

// Generate a random integer64 between min and max
func RandomInt64(min, max int64) int64 {
	return min + globalRand.Int63n(max-min+1)
}

// Generate a random integer32 between min and max
func RandomInt32(min, max int32) int32 {
	return min + globalRand.Int31n(max-min+1)
}

// Generate a random string of length n
func RandomString(n int) string {
	var result strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[globalRand.Intn(k)]
		result.WriteByte(c)
	}

	return result.String()
}

// Generate a random float64 between min and max
func RandomFloat64(min, max int64) float64 {
	return float64(min+globalRand.Int63n(max-min+1)) + globalRand.Float64()
}

package rand

import (
	"math/rand"
	"time"
)

const layout = "2006-01-02 15:04:05"

// RandomTimestamp 区间随机 - 时间戳
func RandomTimestamp(min, max time.Time) time.Time {
	sec := rand.Int63n(max.Unix()-min.Unix()) + min.Unix()
	return time.Unix(sec, 0)
}

// RandomTimestampString 区间随机 - 时间戳
func RandomTimestampString(min, max string) time.Time {
	minT, _ := time.Parse(layout, min)
	maxT, _ := time.Parse(layout, max)
	return RandomTimestamp(minT, maxT)
}

// RandomInt 区间随机 - 整型
func RandomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// RandomFloat 区间随机 - 浮点型
func RandomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

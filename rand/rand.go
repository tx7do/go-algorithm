package rand

import (
	"math/rand"
	"time"
)

// RandomTimestamp 区间随机 - 时间戳
func RandomTimestamp(min, max time.Time) time.Time {
	sec := rand.Int63n(max.Unix()-min.Unix()) + min.Unix()
	return time.Unix(sec, 0)
}

// RandomInt 区间随机 - 整型
func RandomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

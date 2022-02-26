package rand

import (
	"fmt"
	"testing"
	"time"
)

func TestRandomInt(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(RandomInt(1, 1000))
	}
}

func TestRandomFloat(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(RandomFloat(1, 1000))
	}
}

func TestRandomTimestamp(t *testing.T) {
	minTime := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	maxTime := time.Date(2021, 12, 30, 0, 0, 0, 0, time.UTC)
	for i := 0; i < 100; i++ {
		fmt.Println(RandomTimestamp(minTime, maxTime).Format(layout))
	}
}

func TestRandomTimestampString(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(RandomTimestampString("2021-01-01 00:00:00", "2021-12-30 00:00:00").Format(layout))
	}
}

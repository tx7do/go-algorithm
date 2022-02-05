package sorting

import (
	"time"
)

// SleepSort 休眠排序
func SleepSort(array Interface, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	channel := make(chan interface{}, 1)

	for i := begin; i <= end; i++ {
		current := array.Get(i).(int)
		go func(data interface{}, i int) {
			n := 0
			switch t := data.(type) {
			case int:
				n = t
				break
			case float64:
				n = int(t)
			case string:
			}
			sleepTime := time.Duration(n) * time.Second / 4
			time.Sleep(sleepTime)
			channel <- data
		}(current, i)
	}

	j := 0
	for i := begin; i <= end; i++ {
		rev := <-channel
		array.Set(j, rev)
		j++
	}
}

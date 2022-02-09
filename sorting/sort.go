package sorting

import "math"

const (
	AscendingOrder  bool = true
	DescendingOrder bool = false
)

const maxUint = ^uint(0)
const maxInt = int(maxUint >> 1)
const minInt = -maxInt - 1

// Interface of sorter
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
	Get(i int) interface{}
	Set(i int, data interface{})
	Part(low, high int) (part Interface)
}

// InterfaceSlice of sorter
type InterfaceSlice []interface{}

func (x InterfaceSlice) Len() int { return len(x) }
func (x InterfaceSlice) Less(i, j int) bool {
	switch t := x[i].(type) {
	case int:
		return t < x[j].(int)
	case float64:
		return t < x[j].(float64)
	case string:
		return t < x[j].(string)
	}
	return false
}
func (x InterfaceSlice) Swap(i, j int)                       { x[i], x[j] = x[j], x[i] }
func (x InterfaceSlice) Get(i int) interface{}               { return x[i] }
func (x InterfaceSlice) Set(i int, data interface{})         { x[i] = data.(int) }
func (x InterfaceSlice) Part(low, high int) (part Interface) { return x[low:high] }

// IntSlice of sorter
type IntSlice []int

func (x IntSlice) Len() int                            { return len(x) }
func (x IntSlice) Less(i, j int) bool                  { return x[i] < x[j] }
func (x IntSlice) Swap(i, j int)                       { x[i], x[j] = x[j], x[i] }
func (x IntSlice) Get(i int) interface{}               { return x[i] }
func (x IntSlice) Set(i int, data interface{})         { x[i] = data.(int) }
func (x IntSlice) Part(low, high int) (part Interface) { return x[low:high] }

// Float64Slice of sorter
type Float64Slice []float64

func (x Float64Slice) Len() int           { return len(x) }
func (x Float64Slice) Less(i, j int) bool { return x[i] < x[j] || (isNaN(x[i]) && !isNaN(x[j])) }
func (x Float64Slice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func isNaN(f float64) bool                { return f != f }

// StringSlice of sorter
type StringSlice []string

func (x StringSlice) Len() int           { return len(x) }
func (x StringSlice) Less(i, j int) bool { return x[i] < x[j] }
func (x StringSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func IsSorted(data Interface) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

func swapRange(data Interface, a, b, n int) {
	for i := 0; i < n; i++ {
		data.Swap(a+i, b+i)
	}
}

func rotate(data Interface, a, m, b int) {
	i := m - a
	j := b - m

	for i != j {
		if i > j {
			swapRange(data, m-i, m, j)
			i -= j
		} else {
			swapRange(data, m-i, m+j-i, i)
			j -= i
		}
	}
	swapRange(data, m-i, m, i)
}

// merge 原地归并
func merge(array Interface, left, mid, right int) {
	if mid-left == 1 {
		i := mid
		j := right
		for i < j {
			h := int(uint(i+j) >> 1)
			if array.Less(h, left) {
				i = h + 1
			} else {
				j = h
			}
		}

		for k := left; k < i-1; k++ {
			array.Swap(k, k+1)
		}
		return
	}

	if right-mid == 1 {
		i := left
		j := mid
		for i < j {
			h := int(uint(i+j) >> 1)
			if !array.Less(mid, h) {
				i = h + 1
			} else {
				j = h
			}
		}

		for k := mid; k > i; k-- {
			array.Swap(k, k-1)
		}
		return
	}

	midIndex := int(uint(left+right) >> 1)
	n := midIndex + mid
	var start, r int
	if mid > midIndex {
		start = n - right
		r = mid
	} else {
		start = left
		r = mid
	}
	p := n - 1

	for start < r {
		c := int(uint(start+r) >> 1)
		if !array.Less(p-c, c) {
			start = c + 1
		} else {
			r = c
		}
	}

	end := n - start
	if start < mid && mid < end {
		rotate(array, start, mid, end)
	}
	if left < start && start < midIndex {
		merge(array, left, start, midIndex)
	}
	if midIndex < end && end < right {
		merge(array, midIndex, end, right)
	}
}

func less(i, j interface{}) bool {
	switch t := i.(type) {
	case int:
		return t < j.(int)
	case float64:
		return t < j.(float64)
	case string:
		return t < j.(string)
	}
	return false
}

func plus(i, j interface{}) int {
	switch t := i.(type) {
	case int:
		return t + j.(int)
	case float64:
		return int(t + j.(float64))
	case string:
		return 0
	}
	return 0
}
func minus(i, j interface{}) int {
	switch t := i.(type) {
	case int:
		return t - j.(int)
	case float64:
		return int(t - j.(float64))
	case string:
		return 0
	}
	return 0
}
func multi(i, j interface{}) int {
	switch t := i.(type) {
	case int:
		return t * j.(int)
	case float64:
		return int(t * j.(float64))
	case string:
		return 0
	}
	return 0
}

func divisionByTwo(n int) int {
	return int(math.Floor(float64(n >> 1)))
}

func multiplyByTwo(n int) int {
	return int(math.Floor(float64(n << 1)))
}

func greatestPowerOfTwoLessThan(n int) int {
	k := 1
	for k < n {
		k = multiplyByTwo(k)
	}
	return divisionByTwo(k)
}

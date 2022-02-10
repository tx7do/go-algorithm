package sorting

// SmoothSort 平滑排序
// @see https://en.wikipedia.org/wiki/Smoothsort
// @see https://kalkicode.com/smoothsort
// @see https://www.keithschwarz.com/smoothsort/
// @see https://en.wikipedia.org/wiki/Leonardo_numbers
// @see https://www.geeksforgeeks.org/leonardo-number/
// @see https://github.com/Alinshans/LCPP/blob/master/Algorithm/Sort/smooth_sort.cc
// @see https://www.programmingalgorithms.com/algorithm/smooth-sort/
func SmoothSort(array IntSlice, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	var h = newLeonardoHeap()
	h.buildHeap(array, begin, length)
	h.inOrder(array, begin, length)
}

// leonardoNumber are similar to the Fibonacci numbers, and defined as :
// L(0) = L(1) = 1
// L(k + 2) = L(k + 1) + L(k) + 1
type leonardoNumber struct {
	Actual    int
	Companion int
}

func (n *leonardoNumber) diff() int {
	return n.Actual - n.Companion
}
func (n *leonardoNumber) getActual() int {
	return n.Actual
}
func (n *leonardoNumber) getCompanion() int {
	return n.Companion
}
func (n *leonardoNumber) up() *leonardoNumber {
	t := n.Actual
	n.Actual = n.Actual + n.Companion + 1
	n.Companion = t
	return n
}
func (n *leonardoNumber) down() *leonardoNumber {
	t := n.Companion
	n.Companion = n.Actual - n.Companion - 1
	n.Actual = t
	return n
}

type leonardoHeap struct {
	number leonardoNumber
	p      int
}

func newLeonardoHeap() *leonardoHeap {
	var h leonardoHeap
	h.number = leonardoNumber{Actual: 1, Companion: 1}
	h.p = 1
	return &h
}

func (h *leonardoHeap) buildHeap(array IntSlice, begin, length int) {
	i := 1
	for ; i < length; h.p++ {
		if h.p%8 == 3 {
			h.sift(array, begin, i-1)
			h.number.up().up()
			h.p >>= 2
		} else if h.p%4 == 1 {
			if i+h.number.getCompanion() < length {
				h.sift(array, begin, i-1)
			} else {
				h.trinkle(array, begin, i-1, h.p)
			}
			for h.p <<= 1; h.number.down().getActual() > 1; h.p <<= 1 {
			}
		}
		i++
	}

	h.trinkle(array, begin, length-1, h.p)
}

func (h *leonardoHeap) inOrder(array IntSlice, begin, length int) {
	for h.p--; length > 1; h.p-- {
		length--
		if h.number.getActual() == 1 {
			for ; (h.p % 2) == 0; h.p >>= 1 {
				h.number.up()
			}
		} else if h.number.getActual() >= 3 {
			if h.p > 0 {
				h.semiTrinkle(array, begin, length-h.number.diff(), h.p)
			}

			h.number.down()
			h.p <<= 1
			h.p++

			h.semiTrinkle(array, begin, length-1, h.p)

			h.number.down()
			h.p <<= 1
			h.p++
		}
	}
}

func (h leonardoHeap) sift(array IntSlice, first, root int) {
	r := 0
	for h.number.getActual() >= 3 {
		left := first + root - h.number.diff()
		right := first + root - 1
		if !array.Less(left, right) {
			r = root - h.number.diff()
		} else {
			r = root - 1
			h.number.down()
		}

		left = first + root
		right = first + r
		if !array.Less(left, right) {
			break
		}
		array.Swap(right, left)

		root = r
		h.number.down()
	}
}

func (h leonardoHeap) trinkle(array IntSlice, first, root int, p int) {
	for p > 0 {
		for ; (p % 2) == 0; p >>= 1 {
			h.number.up()
		}

		p--
		left := first + root
		right := first + root - h.number.getActual()
		if p == 0 || !array.Less(left, right) {
			break
		}

		if h.number.getActual() == 1 {
			array.Swap(left, right)
			root -= h.number.getActual()
		} else if h.number.getActual() >= 3 {
			r1 := root - h.number.diff()
			r2 := root - h.number.getActual()
			left = first + root - 1
			right = first + r1
			if !array.Less(left, right) {
				r1 = root - 1
				p <<= 1
				h.number.down()
			}

			if !array.Less(first+r2, first+r1) {
				array.Swap(first+root, first+r2)
				root = r2
			} else {
				array.Swap(first+root, first+r1)
				root = r1
				h.number.down()
				break
			}
		}
	}
	h.sift(array, first, root)
}

func (h leonardoHeap) semiTrinkle(array IntSlice, first, root, p int) {
	left := first + root - h.number.getCompanion()
	right := first + root
	if !array.Less(left, right) {
		array.Swap(left, right)
		h.trinkle(array, first, root-h.number.getCompanion(), p)
	}
}

package sorting

import "fmt"

// SmoothSort 平滑排序
// @see https://en.wikipedia.org/wiki/Smoothsort
// @see https://kalkicode.com/smoothsort
// @see https://www.keithschwarz.com/smoothsort/
// @see https://en.wikipedia.org/wiki/Leonardo_numbers
// @see https://www.geeksforgeeks.org/leonardo-number/
// @see https://github.com/Alinshans/LCPP/blob/master/Algorithm/Sort/smooth_sort.cc
func SmoothSort(array IntSlice, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	p := 1

	var t leonardoTree
	var l = LeonardoNumber{Actual: 1, Companion: 1}

	for i := 0; i < length; p++ {
		i++
		if p%8 == 3 {
			t.sift(array, begin, i-1, &l)
			l.up()
			l.up()
			p >>= 2
		} else if p%4 == 1 {
			if i+l.Companion < length {
				t.sift(array, begin, i-1, &l)
			} else {
				t.trinkle(array, begin, i-1, p, &l)
			}
			for p <<= 1; l.down() > 1; p <<= 1 {
				fmt.Println("dddd")
			}
		}
	}

	t.trinkle(array, begin, length-1, p, &l)

	for p--; length > 1; p-- {
		if l.get() == 1 {
			for ; (p % 2) == 0; p >>= 1 {
				l.up()
			}
		} else if l.get() >= 3 {
			if p > 0 {
				t.semiTrinkle(array, begin, length-l.diff(), p, &l)
			}
			l.down()
			p <<= 1
			p++
			t.semiTrinkle(array, begin, length-1, p, &l)
			l.down()
			p <<= 1
			p++
		}

		length--
	}
}

// LeonardoNumber are similar to the Fibonacci numbers, and defined as :
// L(0) = L(1) = 1
// L(k + 2) = L(k + 1) + L(k) + 1
type LeonardoNumber struct {
	Actual    int
	Companion int
}

func (n *LeonardoNumber) diff() int {
	return n.Actual - n.Companion
}
func (n *LeonardoNumber) get() int {
	return n.Actual
}
func (n *LeonardoNumber) up() int {
	t := n.Actual
	n.Actual = n.Actual + n.Companion + 1
	n.Companion = t
	return n.Actual
}
func (n *LeonardoNumber) down() int {
	t := n.Companion
	n.Companion = n.Actual - n.Companion - 1
	n.Actual = t
	return n.Actual
}

type leonardoTree struct {
}

func (t leonardoTree) sift(array IntSlice, first, root int, l *LeonardoNumber) {
	r := 0
	for l.get() > 3 {
		if array.Less(first+root-l.diff(), first+root-1) {
			r = root - l.diff()
		} else {
			r = root - 1
			l.down()
		}
		if array.Less(first+root, first+r) {
			break
		}
		array.Swap(first+root, first+r)
		root = r
		l.down()
	}
}

func (t leonardoTree) trinkle(array IntSlice, first, root, p int, l *LeonardoNumber) {
	for p > 0 {
		for ; (p % 2) == 0; p >>= 1 {
			l.up()
		}
		p--
		if p == 0 || array.Less(first+root, first+root-l.get()) {
			break
		}
		if l.get() == 1 {
			array.Swap(first+root, first+root-l.get())
			root -= l.get()
		} else if l.get() >= 3 {
			r1 := root - l.diff()
			r2 := root - l.get()
			if array.Less(first+root-1, first+r1) {
				r1 = root - 1
				p <<= 1
				l.down()
			}
			if array.Less(first+r2, first+r1) {
				array.Swap(first+root, first+r2)
				root = r2
			} else {
				array.Swap(first+root, first+r1)
				root = r1
				l.down()
				break
			}
		}
	}
	t.sift(array, first, root, l)
}

func (t leonardoTree) semiTrinkle(array IntSlice, first, root, p int, l *LeonardoNumber) {
	left := first + root - l.Companion
	right := first + root
	if array.Less(left, right) {
		array.Swap(left, right)
		t.trinkle(array, first, root-l.Companion, p, l)
	}
}

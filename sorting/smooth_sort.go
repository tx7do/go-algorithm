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

	var t leonardoTree
	t.number = LeonardoNumber{Actual: 1, Companion: 1}

	p := 1
	p = t.buildTree(array, begin, length, p)

	fmt.Println("after======", p, t.number)

	for p--; length > 1; p-- {
		if t.number.getActual() == 1 {
			for ; (p % 2) == 0; p >>= 1 {
				t.number.up()
				fmt.Println("semi heap for ------>>>", p, t.number)
			}
		} else if t.number.getActual() >= 3 {
			if p > 0 {
				t.semiTrinkle(array, begin, length-t.number.diff(), p)
			}

			t.number.down()
			p <<= 1
			p++

			t.semiTrinkle(array, begin, length-1, p)

			t.number.down()
			p <<= 1
			p++
		}

		length--
	}
}

func (t *leonardoTree) buildTree(array IntSlice, begin int, length int, p int) int {
	i := 1
	for ; i < length; p++ {
		fmt.Println("p", i, p, t.number)
		if p%8 == 3 {
			t.sift(array, begin, i-1)
			t.number.up().up()
			p >>= 2
			fmt.Println("p8", p, t.number)
		} else if p%4 == 1 {
			if i+t.number.getCompanion() < length {
				t.sift(array, begin, i-1)
				fmt.Println("p4 sift", p, t.number)
			} else {
				t.trinkle(array, begin, i-1, p)
				fmt.Println("p4 trinkle", p, t.number)
			}
			for p <<= 1; ; p <<= 1 {
				t.number.down()
				fmt.Println("build heap for ------>>>", p, t.number)
				if t.number.getActual() <= 1 {
					break
				}
			}
			fmt.Println("p4", p, t.number)
		}
		i++
	}

	t.trinkle(array, begin, length-1, p)
	return p
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
func (n *LeonardoNumber) getActual() int {
	return n.Actual
}
func (n *LeonardoNumber) getCompanion() int {
	return n.Companion
}
func (n *LeonardoNumber) up() *LeonardoNumber {
	t := n.Actual
	n.Actual = n.Actual + n.Companion + 1
	n.Companion = t
	return n
}
func (n *LeonardoNumber) down() *LeonardoNumber {
	t := n.Companion
	n.Companion = n.Actual - n.Companion - 1
	n.Actual = t
	return n
}

type leonardoTree struct {
	number LeonardoNumber
}

func (t leonardoTree) sift(array IntSlice, first, root int) {
	fmt.Println("sift before", root, t.number)
	r := 0
	for t.number.getActual() >= 3 {
		left := first + root - t.number.diff()
		right := first + root - 1
		if !array.Less(left, right) {
			r = root - t.number.diff()
		} else {
			r = root - 1
			t.number.down()
			fmt.Println("sift <<<<==== 1", first, root, t.number)
		}

		left = first + root
		right = first + r
		if !array.Less(left, right) {
			break
		}
		array.Swap(right, left)

		root = r
		t.number.down()
		fmt.Println("sift <<<<==== 2", first, root, t.number)
	}
}

func (t leonardoTree) trinkle(array IntSlice, first, root int, p int) {
	fmt.Println("trinkle before", first, root, p, t.number)
	for p > 0 {
		for ; (p % 2) == 0; p >>= 1 {
			t.number.up()
			fmt.Println("trinkle->>>>>", p, t.number)
		}

		p--
		left := first + root
		right := first + root - t.number.getActual()
		if p == 0 || !array.Less(left, right) {
			break
		}

		fmt.Println("trinkle swap---> ", left, right)

		if t.number.getActual() == 1 {
			array.Swap(left, right)
			root -= t.number.getActual()
		} else if t.number.getActual() >= 3 {
			r1 := root - t.number.diff()
			r2 := root - t.number.getActual()
			left = first + root - 1
			right = first + r1
			if !array.Less(left, right) {
				r1 = root - 1
				p <<= 1
				t.number.down()
			}

			if !array.Less(first+r2, first+r1) {
				array.Swap(first+root, first+r2)
				root = r2
			} else {
				array.Swap(first+root, first+r1)
				root = r1
				t.number.down()
				break
			}
		}
	}
	//fmt.Println(array, first, root, number)
	t.sift(array, first, root)

	fmt.Println("trinkle after", first, root, p, t.number)
}

func (t leonardoTree) semiTrinkle(array IntSlice, first, root, p int) {
	fmt.Println("semiTrinkle ", first, root, p, t.number)
	left := first + root - t.number.getCompanion()
	right := first + root
	if array.Less(left, right) {
		array.Swap(left, right)
		t.trinkle(array, first, root-t.number.getCompanion(), p)
	}
}

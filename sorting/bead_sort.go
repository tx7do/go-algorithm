package sorting

type beadSlice []byte

// BeadSort 珠排序
// @see https://en.wikipedia.org/wiki/Bead_sort
// @see https://www.wikiwand.com/en/Bead_sort
// @see https://kukuruku.co/post/bead-sort/
// @see https://rosettacode.org/wiki/Sorting_algorithms/Bead_sort
// @see https://www.cs.auckland.ac.nz/~mjd/misc/BeadSort5.pdf
func BeadSort(array Interface, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	maxValue := array.Get(begin).(int)
	i, j := 0, 0
	for i = begin + 1; i <= end; i++ {
		if Less(maxValue, array.Get(i)) {
			maxValue = array.Get(i).(int)
		}
	}

	beadLen := Multi(maxValue, length)
	beads := make(beadSlice, beadLen)

	for i = begin; i <= end; i++ {
		for j = begin; j < array.Get(i).(int); j++ {
			setBeadData(beads, i, j, maxValue, 1)
		}
	}

	sum := 0
	for j = 0; j < maxValue; j++ {
		sum = 0
		for i = begin; i <= end; i++ {
			sum += getBeadData(beads, i, j, maxValue)
			setBeadData(beads, i, j, maxValue, 0)
		}
		for i = length - sum; i < length; i++ {
			setBeadData(beads, i, j, maxValue, 1)
		}
	}

	for i = begin; i <= end; i++ {
		for j = 0; j < maxValue && getBeadData(beads, i, j, maxValue) == 1; j++ {
		}
		array.Set(i, j)
	}
}

func setBeadData(beads beadSlice, x, y, maxValue, data int) {
	beads[x*maxValue+y] = byte(data)
}

func getBeadData(beads beadSlice, x, y, maxValue int) int {
	return int(beads[x*maxValue+y])
}

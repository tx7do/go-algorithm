package sorting

// PigeonholeSort 鸽巢排序
// @see https://en.wikipedia.org/wiki/Pigeonhole_sort
// @see https://zh.wikipedia.org/wiki/%E9%B8%BD%E5%B7%A2%E6%8E%92%E5%BA%8F
func PigeonholeSort(array IntSlice) {
	length := array.Len()
	if length < 2 {
		return
	}

	i := 0

	minValue, maxValue := array[0], array[0]
	for i = 1; i < length; i++ {
		if array[i] < minValue {
			minValue = array[i]
		}
		if array[i] > maxValue {
			maxValue = array[i]
		}
	}

	holeRange := maxValue - minValue + 1
	holes := make([]IntSlice, holeRange)
	for i = 0; i < length; i++ {
		holes[array[i]-minValue] = append(holes[array[i]-minValue], array[i])
	}

	sortedIndex := 0
	j := 0
	for i = 0; i < holeRange; i++ {
		for j = 0; j < holes[i].Len(); j++ {
			array[sortedIndex] = holes[i][j]
			sortedIndex++
		}
	}
}

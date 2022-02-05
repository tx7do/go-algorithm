package sorting

// ProportionExtendSort 扩展比例排序
// @see https://en.wikipedia.org/wiki/Proportion_extend_sort
// @see https://epubs.siam.org/doi/pdf/10.1137/S0097539798342903
func ProportionExtendSort(array Interface, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}
}

package sorting

// TournamentSort 锦标赛排序
// @see https://en.wikipedia.org/wiki/Tournament_sort
// @see http://stepanovpapers.com/TournamentTrees.pdf
// @see https://ocw.snu.ac.kr/sites/default/files/NOTE/490.pdf
func TournamentSort(array Interface, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}
}

type tournamentNode struct {
	data  interface{}
	left  *tournamentNode
	right *tournamentNode
}

type tournamentTree struct {
	root *tournamentNode
}

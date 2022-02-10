package sorting

import (
	"fmt"
	"io"
	"os"
)

// CartesianTreeSort 笛卡尔树排序
// @see https://zh.wikipedia.org/wiki/笛卡尔树
// @see https://oi-wiki.org/ds/cartesian-tree
// @see https://oi-wiki.org/ds/treap/
// @see https://en.wikipedia.org/wiki/Adaptive_sort
// @see http://11011110.livejournal.com/283412.html
// @see http://gradbot.blogspot.in/2010/06/cartesian-tree-sort.html
// @see http://www.keithschwarz.com/interesting/code/?dir=cartesian-tree-sort
// @see https://en.wikipedia.org/wiki/Cartesian_tree#Application_in_sorting
// @see https://ethz.ch/content/dam/ethz/special-interest/infk/chair-program-method/pm/documents/Verify%20This/Challenges%202019/cartesian_trees.pdf
// @see https://people.csail.mit.edu/jshun/alenex2011.pdf
// @see https://medium.com/swlh/cartesian-sequence-and-binary-tree-538cbd0b0ca8
// @see https://github.com/iiitu-force/Hacktoberfest-2021/blob/main/CartesianTreeSort.cpp
func CartesianTreeSort(array IntSlice, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	tree := &cartesianTree{}
	tree.buildTree(array, begin, end)

	tree.print(os.Stdout)

	tree.storeSorted(array, begin, end)
}

type cartesianNode struct {
	data   int
	left   *cartesianNode
	right  *cartesianNode
	parent *cartesianNode
}

type cartesianTree struct {
	root *cartesianNode
	last *cartesianNode
}

// storeSorted Stores inorder traversal of the Cartesian Tree
func (n *cartesianNode) storeSorted(array IntSlice, i *int) {
	if n == nil {
		return
	}

	array[*i] = n.data
	*i++

	n.left.storeSorted(array, i)

	n.right.storeSorted(array, i)
}

func (n *cartesianNode) findRightMost() *cartesianNode {
	curr := n
	for curr != nil {
		if curr.right == nil {
			return curr
		}
		curr = curr.right
	}
	return nil
}

func (n *cartesianNode) findLowestNode(data int) *cartesianNode {
	if n == nil {
		return nil
	}
	if n.data < data {
		return n
	} else if n.parent != nil {
		return n.parent.findLowestNode(data)
	} else {
		return nil
	}
}

// print the Cartesian Tree node
func (n *cartesianNode) print(w io.Writer, ns int, ch rune) {
	if n == nil {
		return
	}

	for i := 0; i < ns; i++ {
		_, _ = fmt.Fprint(w, " ")
	}
	_, _ = fmt.Fprintf(w, "%c:%v\n", ch, n.data)
	n.left.print(w, ns+2, 'L')
	n.right.print(w, ns+2, 'R')
}

// insert a cartesian tree Node
func (n *cartesianNode) insert(tree *cartesianTree, data int) *cartesianNode {
	newNode := &cartesianNode{data: data, left: nil, right: nil, parent: nil}

	lowestNode := tree.last.findLowestNode(data)
	if lowestNode == nil {
		newNode.left = tree.root
		tree.root.parent = newNode
		tree.root = newNode
	} else {
		newNode.left = lowestNode.right
		lowestNode.right = newNode
		newNode.parent = lowestNode
	}
	return newNode
}

func (t *cartesianTree) print(w io.Writer) {
	t.root.print(w, 0, 'M')
}

func (t *cartesianTree) storeSorted(array IntSlice, begin, end int) {
	i := begin
	t.root.storeSorted(array, &i)
}

// buildTree a cartesian tree
func (t *cartesianTree) buildTree(array IntSlice, begin, end int) *cartesianTree {
	for i := begin; i <= end; i++ {
		t.insert(array[i])
	}
	return t
}

// insert a cartesian tree Node
func (t *cartesianTree) insert(data int) *cartesianTree {
	if t.root == nil {
		t.root = &cartesianNode{data: data, left: nil, right: nil, parent: nil}
		t.last = t.root
	} else {
		t.last = t.root.insert(t, data)
	}
	return t
}

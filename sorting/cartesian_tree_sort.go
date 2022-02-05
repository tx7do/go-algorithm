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

func (t *cartesianTree) getRoot() *cartesianNode {
	return t.root
}

func (t *cartesianTree) print(w io.Writer) {
	printCartesianTree(w, t.getRoot(), 0, 'M')
}

func (t *cartesianTree) storeSorted(array IntSlice, begin, end int) {
	i := begin
	storeSortedCartesianTree(t.root, array, &i)
}

func (t *cartesianTree) findLowestNode(node *cartesianNode, data int) *cartesianNode {
	if node.data < data {
		return node
	} else if node.parent != nil {
		return t.findLowestNode(node.parent, data)
	} else {
		return nil
	}
}

// buildTree a cartesian tree
func (t *cartesianTree) buildTree(array IntSlice, begin, end int) *cartesianTree {
	for i := begin; i <= end; i++ {
		t.insert(array[i])
	}
	return t
}

// insert a cartesian tree Node into tree
func (t *cartesianTree) insert(data int) *cartesianTree {
	if t.root == nil {
		t.root = &cartesianNode{data: data, left: nil, right: nil, parent: nil}
		t.last = t.root
	} else {
		t.last = t.root.insert(t, data)
	}
	return t
}

// insert a cartesian tree Node into Node
func (n *cartesianNode) insert(tree *cartesianTree, data int) *cartesianNode {
	newNode := &cartesianNode{data: data, left: nil, right: nil, parent: nil}

	lowestNode := tree.findLowestNode(tree.last, data)
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

// storeSortedCartesianTree Stores inorder traversal of the Cartesian Tree
func storeSortedCartesianTree(node *cartesianNode, array IntSlice, i *int) {
	if node == nil {
		return
	}

	array[*i] = node.data
	*i++

	storeSortedCartesianTree(node.left, array, i)

	storeSortedCartesianTree(node.right, array, i)
}

// printCartesianTree print the Cartesian Tree
func printCartesianTree(w io.Writer, node *cartesianNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		_, _ = fmt.Fprint(w, " ")
	}
	_, _ = fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	printCartesianTree(w, node.left, ns+2, 'L')
	printCartesianTree(w, node.right, ns+2, 'R')
}

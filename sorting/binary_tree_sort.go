package sorting

import (
	"fmt"
	"io"
)

// BinaryTreeSort 二叉树排序
// @see https://en.wikipedia.org/wiki/Tree_sort
func BinaryTreeSort(array Interface, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	tree := &binaryTree{}

	for i := begin; i <= end; i++ {
		tree.insert(array.Get(i))
	}

	tree.inOrder(array)
}

type binaryNode struct {
	data  interface{}
	left  *binaryNode
	right *binaryNode
}

type binaryTree struct {
	root *binaryNode
}

func (t *binaryTree) inOrder(array Interface) {
	if t.root != nil {
		i := 0
		t.root.inOrder(array, &i)
	}
}

// print the binary tree
func (t *binaryTree) print(w io.Writer) {
	t.root.print(w, 0, 'M')
}

// insert a binary tree Node into tree
func (t *binaryTree) insert(data interface{}) *binaryTree {
	if t.root == nil {
		t.root = &binaryNode{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

// insert a binary tree Node into Node
func (n *binaryNode) insert(data interface{}) {
	if n == nil {
		return
	} else if Less(data, n.data) {
		if n.left == nil {
			n.left = &binaryNode{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &binaryNode{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}

// inOrder traversal of the BST
func (n *binaryNode) inOrder(array Interface, i *int) {
	if n == nil {
		return
	}

	if n.left != nil {
		n.left.inOrder(array, i)
	}

	array.Set(*i, n.data)
	*i++

	if n.right != nil {
		n.right.inOrder(array, i)
	}
}

// print the BST
func (n *binaryNode) print(w io.Writer, ns int, ch rune) {
	if n == nil {
		return
	}

	for i := 0; i < ns; i++ {
		_, _ = fmt.Fprint(w, " ")
	}
	_, _ = fmt.Fprintf(w, "%c:%v\n", ch, n.data)

	if n.left != nil {
		n.left.print(w, ns+2, 'L')
	}
	if n.right != nil {
		n.right.print(w, ns+2, 'R')
	}
}

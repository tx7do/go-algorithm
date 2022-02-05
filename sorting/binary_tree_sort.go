package sorting

import (
	"fmt"
	"io"
)

// BinaryTreeSort 二叉树排序
// @see https://en.wikipedia.org/wiki/Tree_sort
func BinaryTreeSort(array IntSlice) {
	length := array.Len()
	if length < 2 {
		return
	}

	i := 0

	// Construct the BST
	tree := &binaryTree{}
	for ; i < length; i++ {
		tree.insert(array[i])
	}

	// Store inorder traversal of the BST
	i = 0
	storeSorted(tree.root, array, &i)
}

type binaryNode struct {
	left  *binaryNode
	right *binaryNode
	data  int
}

type binaryTree struct {
	root *binaryNode
}

// insert a binary node into tree
func (t *binaryTree) insert(data int) *binaryTree {
	if t.root == nil {
		t.root = &binaryNode{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

// insert a binary node into node
func (n *binaryNode) insert(data int) {
	if n == nil {
		return
	} else if data <= n.data {
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

// storeSorted Stores inorder traversal of the BST
func storeSorted(node *binaryNode, array IntSlice, i *int) {
	if node == nil {
		return
	}

	storeSorted(node.left, array, i)
	array[*i] = node.data
	*i++
	storeSorted(node.right, array, i)
}

// printTree print the BST
func printTree(w io.Writer, node *binaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		_, _ = fmt.Fprint(w, " ")
	}
	_, _ = fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	printTree(w, node.left, ns+2, 'L')
	printTree(w, node.right, ns+2, 'R')
}

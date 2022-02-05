package sorting

import (
	"fmt"
	"io"
)

// SplaySort 伸展树排序
// @see https://en.wikipedia.org/wiki/Splaysort
// @see https://en.wikipedia.org/wiki/Splay_tree
// @see https://11011110.github.io/blog/2014/01/21/splaysort-versus-cartesian.html
// @see https://kalkicode.com/splay-tree-deletion-in-go
func SplaySort(array Interface, begin, end int) {
	length := end - begin + 1
	if length < 2 {
		return
	}

	tree := &splayTree{}
	for i := begin; i <= end; i++ {
		tree.insert(array.Get(i))
	}

	tree.inOrder(array)
}

type splayNode struct {
	data   interface{}
	left   *splayNode
	right  *splayNode
	parent *splayNode
}

type splayTree struct {
	root *splayNode
}

// print the Splay Tree
func (t *splayTree) print(w io.Writer) {
	t.root.print(w, 0, 'M')
}

func (t *splayTree) insert(data interface{}) *splayTree {
	node := &splayNode{data: data, left: nil, right: nil}

	if t.root == nil {
		t.root = node
	} else {
		temp := t.root

		for temp != nil {
			if Less(temp.data, data) {
				if temp.right == nil {
					temp.right = node
					node.parent = temp
					temp = nil
				} else {
					temp = temp.right
				}
			} else {
				if temp.left == nil {
					temp.left = node
					node.parent = temp
					temp = nil
				} else {
					temp = temp.left
				}
			}
		}

		node.applyRotation()
	}
	t.root = node

	return t
}

// print the Splay Tree node
func (n *splayNode) print(w io.Writer, ns int, ch rune) {
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

// zig rotation
func (n *splayNode) zig() {
	var parent = n.parent
	parent.left = n.right
	if n.right != nil {
		n.right.parent = parent
	}
	n.right = parent
	n.parent = parent.parent
	parent.parent = n
}

// zag rotation
func (n *splayNode) zag() {
	var parent = n.parent
	parent.right = n.left
	if parent.right != nil {
		parent.right.parent = parent
	}
	n.left = parent
	n.parent = parent.parent
	parent.parent = n
}

// zig zig rotation
func (n *splayNode) zigZig() {
	var parent = n.parent
	var grandParent = n.parent.parent
	parent.left = n.right
	if n.right != nil {
		n.right.parent = parent
	}
	n.right = parent
	parent.parent = n
	grandParent.left = parent.right
	if parent.right != nil {
		parent.right.parent = grandParent
	}
	parent.right = grandParent
	n.parent = grandParent.parent
	if grandParent.parent != nil {
		if grandParent.parent.left != nil &&
			grandParent.parent.left == grandParent {
			grandParent.parent.left = n
		} else {
			grandParent.parent.right = n
		}
	}
	grandParent.parent = parent
}

// zag zag rotation
func (n *splayNode) zagZag() {
	var parent = n.parent
	var grandParent = n.parent.parent
	parent.right = n.left
	if n.left != nil {
		n.left.parent = parent
	}
	n.left = parent
	n.parent = grandParent.parent
	if grandParent.parent != nil {
		if grandParent.parent.left != nil &&
			grandParent.parent.left == grandParent {
			grandParent.parent.left = n
		} else {
			grandParent.parent.right = n
		}
	}
	parent.parent = n
	grandParent.right = parent.left
	if parent.left != nil {
		parent.left.parent = grandParent
	}
	parent.left = grandParent
	grandParent.parent = parent
}

// Zag zig rotation
func (n *splayNode) zagZig() {
	var parent = n.parent
	var grandParent = n.parent.parent
	parent.left = n.right
	if n.right != nil {
		n.right.parent = parent
	}
	grandParent.right = n
	n.parent = grandParent
	n.right = parent
	parent.parent = n
}

// Zig zag rotation
func (n *splayNode) zigZag() {
	var parent = n.parent
	var grandParent = n.parent.parent
	parent.right = n.left
	if n.left != nil {
		n.left.parent = parent
	}
	grandParent.left = n
	n.parent = grandParent
	n.left = parent
	parent.parent = n
}

func (n *splayNode) applyRotation() {
	if n.parent != nil {
		if n.parent.left != nil &&
			n.parent.left == n &&
			n.parent.parent == nil {
			n.zig()
		} else if n.parent.right != nil &&
			n.parent.right == n &&
			n.parent.parent == nil {
			n.zag()
		} else if n.parent.left != nil &&
			n.parent.left == n &&
			n.parent.parent.left != nil &&
			n.parent.parent.left == n.parent {
			n.zigZig()
		} else if n.parent.right != nil &&
			n.parent.right == n &&
			n.parent.parent.right != nil &&
			n.parent.parent.right == n.parent {
			n.zagZag()
		} else if n.parent.right != nil &&
			n.parent.right == n &&
			n.parent.parent != nil &&
			n.parent.parent.left != nil &&
			n.parent.parent.left == n.parent {
			n.zigZag()
		} else if n.parent.left != nil &&
			n.parent.left == n &&
			n.parent.parent != nil &&
			n.parent.parent.right != nil &&
			n.parent.parent.right == n.parent {
			n.zagZig()
		} else {
			return
		}
		n.applyRotation()
	}
}

func (t *splayTree) preOrder(array Interface) {
	if t.root != nil {
		i := 0
		t.root.preOrder(array, &i)
	}
}
func (t *splayTree) inOrder(array Interface) {
	if t.root != nil {
		i := 0
		t.root.inOrder(array, &i)
	}
}
func (t *splayTree) postOrder(array Interface) {
	if t.root != nil {
		i := 0
		t.root.postOrder(array, &i)
	}
}

func (n *splayNode) preOrder(array Interface, i *int) {
	if n == nil {
		return
	}
	array.Set(*i, n.data)
	*i++

	if n.left != nil {
		n.left.preOrder(array, i)
	}
	if n.right != nil {
		n.right.preOrder(array, i)
	}
}

func (n *splayNode) inOrder(array Interface, i *int) {
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

func (n *splayNode) postOrder(array Interface, i *int) {
	if n == nil {
		return
	}
	if n.left != nil {
		n.left.postOrder(array, i)
	}
	if n.right != nil {
		n.right.postOrder(array, i)
	}
	array.Set(*i, n.data)
	*i++
}

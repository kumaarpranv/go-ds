package bst

import "fmt"

type Node struct {
	val   uint32
	left  *Node
	right *Node
}

type BSTree struct {
	Root *Node
}

func NewNode(val uint32) *Node {
	var n *Node = &Node{
		val: val,
	}
	return n
}

func NewBSTree(val uint32) *BSTree {
	var root *Node = NewNode(val)
	var m *BSTree = &BSTree{
		Root: root,
	}
	return m
}

func InsertNode(node *Node, val uint32) {
	var queue = make([]*Node, 0)
	queue = append(queue, node)
	for len(queue) > 0 {
		el := queue[0]
		queue = queue[1:]
		if el.left == nil {
			el.left = NewNode(val)
			break
		} else if el.right == nil {
			el.right = NewNode(val)
			break
		} else {
			queue = append(queue, el.left)
			queue = append(queue, el.right)
		}
	}
}

func (b *BSTree) BSTInsert(val uint32) {
	var queue = make([]*Node, 0)
	queue = append(queue, b.Root)
	for len(queue) > 0 {
		el := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if val <= el.val {
			if el.left == nil {
				el.left = NewNode(val)
			} else {
				queue = append(queue, el.left)
			}
		} else if val > el.val {
			if el.right == nil {
				el.right = NewNode(val)
			} else {
				queue = append(queue, el.right)
			}
		}
	}
}

func (b *BSTree) BSTPreorder() {
	var queue = make([]*Node, 0)
	queue = append(queue, b.Root)
	for len(queue) > 0 {
		el := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		fmt.Println(el.val)

		if el.right != nil {
			queue = append(queue, el.right)
		}
		if el.left != nil {
			queue = append(queue, el.left)
		}
	}
}

func (b *BSTree) BSTPostorder() {
	var queue = make([]*Node, 0)
	var res = make([]uint32, 0)
	queue = append(queue, b.Root)
	for len(queue) > 0 {
		el := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		res = append(res, el.val)
		if el.left != nil {
			queue = append(queue, el.left)
		}
		if el.right != nil {
			queue = append(queue, el.right)
		}
	}

	for i := len(res) - 1; i >= 0; i-- {
		fmt.Println(res[i])
	}
}

func (b *BSTree) BSTInorder() {
	var queue = make([]*Node, 0)
	el := b.Root
	for el != nil || len(queue) > 0 {
		for el != nil {
			queue = append(queue, el)
			el = el.left
		}
		el = queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		fmt.Println(el.val)
		el = el.right
	}
}

func PrintTree(node *Node) {
	var queue = make([]*Node, 0)
	queue = append(queue, node)
	for len(queue) > 0 {
		el := queue[0]
		queue = queue[1:]
		fmt.Println(el.val)
		if el.left != nil {
			queue = append(queue, el.left)
		}
		if el.right != nil {
			queue = append(queue, el.right)
		}
	}

}

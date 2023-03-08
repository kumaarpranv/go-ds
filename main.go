package main

import "fmt"

type Node struct {
	val   uint32
	left  *Node
	right *Node
}

type BSTree struct {
	root *Node
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
		root: root,
	}
	return m
}

func InsertNode(node *Node, val uint32, queue []*Node) {
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

func (b *BSTree) BSTInsert(val uint32, queue []*Node) {
	queue = append(queue, b.root)
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

func PrintTree(node *Node, queue []*Node) {
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
func main() {
	var minh *BSTree = NewBSTree(6)
	var queue = make([]*Node, 0)
	for i := 2; i < 11; i = i + 1 {
		minh.BSTInsert(uint32(i), queue)
	}
	PrintTree(minh.root, queue)

}

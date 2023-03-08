package main

import "fmt"

type Node struct {
	val   uint32
	left  *Node
	right *Node
}

func NewNode(val uint32) *Node {
	var n *Node = &Node{
		val: val,
	}
	return n
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

func PrintTree(node *Node, queue []*Node) {
	queue = append(queue, node)
	for len(queue) > 0 {
		el := queue[0]
		queue = queue[1:]
		fmt.Println(el.val)
		if el.left != nil {
			queue = append(queue, el.left)
			fmt.Println("~", el.left.val)
		}
		if el.right != nil {
			queue = append(queue, el.right)
			fmt.Println("-", el.right.val)
		}
	}

}
func main() {
	var root *Node = NewNode(1)
	var queue = make([]*Node, 0)
	curr := root
	queue = append(queue)
	for i := 2; i < 11; i = i + 1 {
		InsertNode(root, uint32(i), queue)
	}
	PrintTree(curr, queue)

}

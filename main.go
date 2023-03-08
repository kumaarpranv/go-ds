package main

import (
	. "github.com/kumaarpranv/go-ds/bst"
)

func main() {
	var minh *BSTree = NewBSTree(6)
	var queue = make([]*Node, 0)
	for i := 2; i < 11; i = i + 1 {
		minh.BSTInsert(uint32(i), queue)
	}
	PrintTree(minh.Root, queue)

}

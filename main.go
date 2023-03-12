package main

import (
  // . "github.com/kumaarpranv/go-ds/bst"
  . "github.com/kumaarpranv/go-ds/linkedlist"
)

func main() {
    // var minh *BSTree = NewBSTree(6)
    // for i := 2; i < 11; i = i + 1 {
    //   minh.BSTInsert(uint32(i))
    // }
    //PrintTree(minh.Root)
    // minh.BSTPreorder()
    var ll *LinkedList = NewLinkedList(1);
    ll.Append(2);ll.Append(3);ll.Append(4);
    ll.PrintList();
    ll.RemoveTail(); ll.RemoveTail();
    ll.PrintList();
}

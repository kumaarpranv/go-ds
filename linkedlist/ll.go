package linkedlist

import "fmt"

type ListNode struct {
  val uint32
  next *ListNode
}

type LinkedList struct {
  Head *ListNode
  Tail *ListNode
  Length uint32
}

func NewListNode(val uint32) (*ListNode) {
  var head *ListNode = &ListNode{
      val: val,
      next: nil,
  };
  return head;
}

func NewLinkedList(val uint32) (*LinkedList) {
  var tmp *ListNode = NewListNode(val);
  var ll *LinkedList = &LinkedList{
    Head: tmp,
    Tail: tmp,
    Length: 1,
  } 
  return ll;
}

func (ll *LinkedList) Append(val uint32) {
 var tmp *ListNode = ll.Tail;
 var tail *ListNode = NewListNode(val);
 tmp.next = tail;
 ll.Tail = tail;
}


func (ll *LinkedList) RemoveTail() {
 var tmp *ListNode = ll.Head;
 for tmp.next != ll.Tail {
   tmp = tmp.next;
 }
 tmp.next = nil;
 ll.Tail = tmp;
}

func (ll *LinkedList) PrintList() {
  var tmp = ll.Head;
  for tmp != nil {
    fmt.Print(tmp.val);
    if tmp.next != nil {
      fmt.Print("-")
    } else {
      fmt.Print("\n");
    }
    tmp = tmp.next;
  }
}

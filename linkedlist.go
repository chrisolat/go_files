package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type LL struct {
	head *Node
}

func (ll *LL) init() {
	ll.head = &Node{
		val:  0,
		next: nil,
	}
}

func (ll *LL) insertHead(val int) {
	node := &Node{
		val:  val,
		next: nil,
	}
	node.next = ll.head
	ll.head = node
}

func (ll *LL) insertTail(val int) {
	cur := ll.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = &Node{
		val:  val,
		next: nil,
	}
}

func (ll *LL) printList() {
	cur := ll.head
	for cur != nil {
		fmt.Println(cur.val)
		cur = cur.next
	}
}

func main() {
	ll := LL{}
	ll.init()
	ll.insertHead(1)
	ll.insertHead(2)
	ll.insertTail(3)
	ll.printList()
}

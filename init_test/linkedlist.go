package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type LL struct {
	head *Node
	size int
}

func New() *LL {
	return &LL{}
}

func (ll *LL) InsertHead(val int) {
	cur := ll.head
	if cur == nil {
		ll.head = &Node{val: val, next: nil}
		return
	}
	node := &Node{val: val, next: nil}
	node.next = cur
	ll.head = node
	return
}

func (ll *LL) InsertTail(val int) {
	cur := ll.head
	if cur == nil {
		ll.head = &Node{val: val, next: nil}
		return
	}
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = &Node{val: val, next: nil}
	return
}

func (ll *LL) PopHead() int {
	if ll.head == nil {
		return -1
	}
	item := ll.head.val
	ll.head = ll.head.next
	return item
}

func (ll *LL) PopTail() int {
	if ll.head == nil {
		return -1
	}
	cur := ll.head
	var prev *Node
	for cur.next != nil {
		prev = cur
		cur = cur.next
	}
	item := -1
	if prev != nil {
		item = cur.val
		prev.next = cur.next
	} else {
		item = ll.head.val
		ll.head = nil
	}
	return item
}

func (ll *LL) PrintLL() {
	cur := ll.head
	for cur != nil {
		fmt.Println(cur.val)
		cur = cur.next
	}
}

func main() {
	ll := New()
	ll.InsertHead(1)
	ll.InsertTail(2)
	ll.PopHead()
	ll.PrintLL()
}

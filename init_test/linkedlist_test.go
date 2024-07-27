package main

import "testing"

func TestInsertHead(t *testing.T) {
	ll := New()
	ll.InsertHead(1)
	ll.InsertHead(2)
	if ll.head == nil || ll.head.val != 2 {
		t.Errorf("InsertHead failed, got %d, want %d", ll.head.val, 2)
	}

	if ll.head.next == nil || ll.head.next.val != 1 {
		t.Errorf("InsertHead failed, got %d, want %d", ll.head.next.val, 1)
	}
}
func TestInsertTail(t *testing.T) {
	ll := New()
	ll.InsertTail(2)
	ll.InsertTail(1)
	ll.InsertTail(0)

	if ll.head == nil || ll.head.val != 2 {
		t.Errorf("InsertTail failed, got %d, want %d", ll.head.val, 2)
	}

	if ll.head.next == nil || ll.head.next.val != 1 {
		t.Errorf("InsertTail failed, got %d, want %d", ll.head.next.val, 1)
	}
}

func TestPopHead(t *testing.T) {
	ll := New()
	ll.InsertHead(1)
	ll.InsertHead(2)

	val := ll.PopHead()
	if val != 2 {
		t.Errorf("PopHead failed, got %d, want %d", val, 2)
	}

	val = ll.PopHead()
	if val != 1 {
		t.Errorf("PopHead failed, got %d, want %d", val, 1)
	}

	val = ll.PopHead()
	if val != -1 {
		t.Errorf("PopHead failed, got %d, want %d", val, -1)
	}
}

func TestPopTail(t *testing.T) {
	ll := New()
	ll.InsertTail(1)
	ll.InsertTail(2)

	val := ll.PopTail()
	if val != 2 {
		t.Errorf("PopHead failed, got %d, want %d", val, 2)
	}

	val = ll.PopTail()
	if val != 1 {
		t.Errorf("PopHead failed, got %d, want %d", val, 1)
	}

	val = ll.PopTail()
	if val != -1 {
		t.Errorf("PopHead failed, got %d, want %d", val, -1)
	}

}
func TestPrintLL(t *testing.T) {
	ll := New()
	ll.InsertTail(1)
	ll.InsertTail(2)
	ll.InsertTail(3)

	expected := []int{1, 2, 3}
	cur := ll.head
	for _, v := range expected {
		if cur == nil || cur.val != v {
			t.Errorf("PrintLL failed, got %d, want %d", cur.val, v)
		}
		cur = cur.next
	}
	if cur != nil {
		t.Error("PrintLL failed, linked list has more elements than expected")
	}
}

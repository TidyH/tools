package linkedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head  *ListNode
	Tail  *ListNode
	Count int
}

// constructor
func NewLinkedList() *LinkedList {
	return &LinkedList{Head: nil, Tail: nil, Count: 0}
}

func (ll *LinkedList) Append(val int) {
	newNode := &ListNode{Val: val, Next: nil}

	if ll.Head == nil {
		ll.Head = newNode
		ll.Tail = newNode
	} else {
		ll.Tail.Next = newNode
		ll.Tail = newNode
	}

	ll.Count++
}

func (ll *LinkedList) Prepend(val int) {
	newNode := &ListNode{Val: val, Next: nil}

	if ll.Head == nil {
		ll.Head = newNode
		ll.Tail = newNode
	} else {
		newNode.Next = ll.Head
		ll.Head = newNode
	}

	ll.Count++
}

func (ll *LinkedList) PeekHead() (int, bool) {
	if ll.Head == nil {
		return 0, false
	}
	return ll.Head.Val, true
}

func (ll *LinkedList) PeekTail() (int, bool) {
	if ll.Tail == nil {
		return 0, false
	}
	return ll.Tail.Val, true
}

func (ll *LinkedList) Len() int {
	return ll.Count
}

func (ll *LinkedList) RemoveFirst() (int, bool) {
	if ll.Head == nil {
		return 0, false
	}
	if ll.Len() == 1 {
		removed := ll.Head
		ll.Head = nil
		ll.Tail = nil
		ll.Count--
		return removed.Val, true
	} else {
		removed := ll.Head
		ll.Head = ll.Head.Next
		ll.Count--
		return removed.Val, true
	}
}

func (ll *LinkedList) RemoveLast() (int, bool) {
	if ll.Head == nil {
		return 0, false
	}

	if ll.Len() == 1 {
		removed := ll.Tail
		ll.Head = nil
		ll.Tail = nil
		return removed.Val, true
	} else {
		var prev *ListNode
		for curr := ll.Head; curr != ll.Tail; curr = curr.Next {
			prev = curr
		}
		removed := ll.Tail
		prev.Next = nil
		ll.Tail = prev
		ll.Count--
		return removed.Val, true
	}
}
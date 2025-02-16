package main

func reverseList(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode = nil

	for curr != nil {
		println(head.Val)
		next := curr.Next // get next node
		curr.Next = prev  // reverse the pointer
		prev = curr       // set previous to current node
		curr = next       // set curr to next node
	}

	return prev
}

func main() {
	println("206. Reverse Linked List")

	node2 := ListNode{
		Val:  2,
		Next: nil,
	}

	node1 := ListNode{
		Val:  1,
		Next: &node2,
	}

	test1 := LinkedList{
		Head:   &node1,
		Length: 2,
	}

	// node3 := ListNode{
	// 	Val:  3,
	// 	Next: *node4,
	// }
	// node4 := ListNode{
	// 	Val:  4,
	// 	Next: nil,
	// }

	reverseList(test1.Head)
}

package main

import (
	"leetcode/internal/linkedlists"
)

func mergeTwoLists(list1 *linkedlists.ListNode, list2 *linkedlists.ListNode) *linkedlists.ListNode {
	dummy := &linkedlists.ListNode{}
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}

		current = current.Next
	}

	if list1 == nil {
		current.Next = list2
	} else if list2 == nil {
		current.Next = list1
	}

	return dummy.Next
}

func main() {
	println("21. Merge two sorted lists")

	list1 := linkedlists.LinkedList{
		Head: &linkedlists.ListNode{
			Val: 1,
			Next: &linkedlists.ListNode{
				Val: 2,
				Next: &linkedlists.ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
		Count: 3,
	}

	list2 := linkedlists.LinkedList{
		Head: &linkedlists.ListNode{
			Val: 1,
			Next: &linkedlists.ListNode{
				Val: 3,
				Next: &linkedlists.ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
		Count: 3,
	}

	// Input: list1 = [1,2,4], list2 = [1,3,4]
	// Output: [1,1,2,3,4,4]
	h := mergeTwoLists(list1.Head, list2.Head)

	for h != nil {
		println(h.Val)
		h = h.Next
	}

}

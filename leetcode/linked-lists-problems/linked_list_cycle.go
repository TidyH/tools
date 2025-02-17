package main

import (
	"fmt"
	"leetcode/internal/linkedlists"
)

// naive approach, space is O(n), time is O(n) time is okay space is not
// func hasCycle(head *linkedlists.ListNode) bool {
// 	seen := make(map[*linkedlists.ListNode]int)

// 	for head != nil {
// 		// println(head, head.Val)

// 		if val, ok := seen[head]; ok {
// 			println(val)
// 			return true
// 		} else {
// 			seen[head]++
// 		}

// 		head = head.Next
// 	}

// 	return false
// }

// fast and slow list implementation
// if fast ends up on the same address as slow, there is a cycle
func hasCycle(head *linkedlists.ListNode) bool {
	fast := head
	slow := head

	for fast != nil && slow != nil {
		if fast.Next == nil || slow.Next == nil {
			return false
		}

		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}
	return false
}


func main(){
	println("141. Linked List Cycle")

	// Input: head = [1,2,3,4], index = 1
	// Output: true
	// println("Creating test linked list")
	tc1 := linkedlists.NewLinkedList()
	// fmt.Print("%v\n", tc1)

	// println("\nAppending new nodes")
	for i := 1; i < 4; i++ {
		tc1.Append(i, nil)
	}
	tc1.Append(4, tc1.Head.Next)
	// fmt.Print("%v\n", tc1)

	// curr := tc1.Head
	// for curr != nil {
	// 	fmt.Printf("%v\n", curr)
	// 	curr = curr.Next
	// }

	fmt.Printf("%v\n", hasCycle(tc1.Head))
}
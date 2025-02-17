package main

import (
	"fmt"
	"leetcode/internal/linkedlists"
)

func hasCycle(head *linkedlists.ListNode) bool {
	// determine if linked list has a cycle
	// in my own words, the linked list does not  have a tail, it will loop forver at some point
	// we just need to return true if we determine this to be true
	// hashmap of pointer and index? (we might not need index? maybe just true?)
	// each pointer should be unique so if we run across it again, we've cycled

	seen := make(map[*linkedlists.ListNode]int)

	for head != nil { // should be infinite in our case for a cycle
		if val, ok := seen[head]; ok {
			if val > 1 {
				return true
			}
		}
	}

	return false
}

func main(){
	println("141. Linked List Cycle")

	// Input: head = [1,2,3,4], index = 1
	// Output: true
	tc1 := linkedlists.NewLinkedList()
	for i := 1; i == 4; i++ {
		tc1.Append(i)
	}

	for tc1 != nil {
		fmt.Print(tc1.Head)
		tc1.Head = tc1.Head.Next
	}

}
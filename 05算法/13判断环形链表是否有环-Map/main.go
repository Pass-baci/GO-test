package main

type ListNode struct {
	Val int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]int)
	for head != nil {
		if _, ok := m[head]; ok {
			return true
		}
		m[head] = 1
		head = head.Next
	}
	return false
}

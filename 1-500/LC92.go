package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
}

func reverseBetween(head *ListNode, left, right int) *ListNode {
	dummyNode := &ListNode{Val: -1, Next: head}
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}
	leftNode := pre.Next
	curr := rightNode.Next

	pre.Next = nil
	rightNode.Next = nil

	reverse(leftNode)

	pre.Next = rightNode
	leftNode.Next = curr
	return dummyNode.Next
}

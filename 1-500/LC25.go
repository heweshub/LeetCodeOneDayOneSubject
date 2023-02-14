package main

// 反转链表
func reverse1(head, tail *ListNode) (*ListNode, *ListNode) {
	// 尾结点之后的结点连接到这段链表的开头
	prev := tail.Next
	// 下面的逻辑和反转链表一致
	cur := head
	for prev != tail {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	// 返回新的头尾结点
	return tail, head
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummyHead := &ListNode{Next: head}
	prev := dummyHead
	for head != nil {
		tail := prev
		// 遍历之后的k个结点，找到这段链表的头尾结点
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummyHead.Next
			}
		}
		next := tail.Next
		head, tail = reverse1(head, tail)
		// 再拼接起来
		prev.Next = head
		tail.Next = next
		prev = tail
		head = tail.Next
	}
	return dummyHead.Next
}

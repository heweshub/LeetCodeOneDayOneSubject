package main

func merge(h1, h2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	tmp, tmp1, tmp2 := dummyHead, h1, h2
	for tmp1 != nil && tmp2 != nil {
		if tmp1.Val < tmp2.Val {
			tmp.Next = tmp1
			tmp1 = tmp1.Next
		} else {
			tmp.Next = tmp2
			tmp2 = tmp2.Next
		}
		tmp = tmp.Next
	}
	if tmp1 != nil {
		tmp.Next = tmp1
	} else {
		tmp.Next = tmp2
	}
	return dummyHead.Next
}

func sort1(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = tail
		return head
	}
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	return merge(sort1(head, slow), sort1(slow, tail))
}

func sortList(head *ListNode) *ListNode {
	return sort1(head, nil)
}

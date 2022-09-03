package main

func mergeNodes(head *ListNode) *ListNode {
    ans := head
    for node, sum := head.Next, 0;node != nil;node = node.Next {
        if node.Val > 0 {
            sum += node.Val
        } else {
            ans.Next.Val = sum
            ans = ans.Next
            sum = 0
        }
    }
    ans.Next = nil
    return head.Next
}

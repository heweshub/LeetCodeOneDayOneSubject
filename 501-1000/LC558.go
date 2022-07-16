package main

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func intersect(q1, q2 *Node) *Node {
	if q1.IsLeaf {
		if q1.Val {
			return &Node{Val: true, IsLeaf: true}
		}
		return q2
	}
	if q2.IsLeaf {
		return intersect(q2, q1)
	}
	o1 := intersect(q1.TopLeft, q2.TopLeft)
	o2 := intersect(q1.TopRight, q2.TopRight)
	o3 := intersect(q1.BottomLeft, q2.BottomLeft)
	o4 := intersect(q1.BottomRight, q2.BottomRight)
	if o1.IsLeaf && o2.IsLeaf && o3.IsLeaf && o4.IsLeaf && o1.Val == o2.Val && o1.Val == o3.Val && o1.Val == o4.Val {
		return &Node{Val: o1.Val, IsLeaf: true}
	}
	return &Node{false, false, o1, o2, o3, o4}
}

package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := map[*Node]*Node{}
	var cg func(*Node) *Node
	cg = func(node *Node) *Node {
		if node == nil {
			return node
		}
		if _, ok := visited[node]; ok {
			return visited[node]
		}
		cloneNode := &Node{node.Val, []*Node{}}

		visited[node] = cloneNode

		for _, n := range node.Neighbors {
			cloneNode.Neighbors = append(cloneNode.Neighbors, cg(n))
		}
		return cloneNode
	}
	return cg(node)
}

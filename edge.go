package main

type Edge struct {
	Value byte
	Next  *Node
}

func (edge *Edge) addNode(node *Node) {
	edge.Next = node
}

func (edge *Edge) nextNode() *Node {
	if edge.Next == nil {
		node := newNode()
		edge.addNode(&node)
	}
	return edge.Next
}

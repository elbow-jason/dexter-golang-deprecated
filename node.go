package main

type Node struct {
	Edges []*Edge
	ID    int
}

func newEmptyEdges() []*Edge {
	return make([]*Edge, 255)
}

func (node *Node) hasValue(value byte) bool {
	return (node.Edges[value] != nil)
}

func (node *Node) getEdge(value byte) *Edge {
	return node.Edges[value]
}

func (node *Node) hasEdge(edge Edge) bool {
	return node.hasValue(edge.Value)
}

func (node *Node) addEdge(edge *Edge) {
	node.Edges[edge.Value] = edge
}

var count int = 0

func countNext() int {
	count = count + 1
	return count
}

func newNode() Node {
	return Node{
		ID:    countNext(),
		Edges: newEmptyEdges(),
	}
}

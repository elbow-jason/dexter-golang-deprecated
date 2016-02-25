package main

// Node is a type that has an ID and edges that
// are indexed by values between 0 and 255
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

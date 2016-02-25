package dexter

// Edge has a byte and points to the next node
// in the graph/FST.
type Edge struct {
	Value byte
	Next  *Node
}

func (edge *Edge) addNode(node *Node) {
	edge.Next = node
}

func (edge *Edge) nextNode(fst *FST) *Node {
	if edge.Next == nil {
		node := fst.newNode()
		edge.addNode(&node)
	}
	return edge.Next
}

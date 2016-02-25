package dexter

import (
	"testing"
)

func TestEdgeCanAddNode(t *testing.T) {
	fst := NewFST()
	node := fst.newNode()
	edge := Edge{}
	edge.addNode(&node)
	if &node != edge.Next {
		t.Fatal("Expected edge.Next to be the same node that was passed to addNode")
	}
}

package dexter

import (
	"errors"
)

// FST - A finite state transducer that holds a searchable graph
type FST struct {
	Start   *Node
	Finish  *Node
	counter int
}

var errNotFound = errors.New("Not Found")

func NewFST() FST {
	fst := FST{}
	fst.initialize()
	return fst
}

// AddString is used to add a branch to the FST.
func (fst *FST) AddString(word string) {
	var (
		outterEdge *Edge
		node       = fst.Start
	)
	edges := stringToEdges(word)
	for _, edge := range edges {
		if !node.hasValue(edge.Value) {
			node.addEdge(edge)
		} else {
			edge = node.getEdge(edge.Value)
		}
		node = edge.getOrCreateNode(fst)
		outterEdge = edge
	}
	outterEdge.Next = fst.Finish
}

// Search for strings within an FST's graphs
func (fst *FST) Search(word string) error {
	var (
		edge *Edge
		node = fst.Start
	)
	for _, value := range []byte(word) {
		if node.hasValue(value) {
			edge = node.Edges[value]
			node = edge.getOrCreateNode(fst)
		} else {
			return errNotFound
		}
	}
	return nil
}

func stringToEdges(word string) []*Edge {
	var edges []*Edge
	for _, value := range []byte(word) {
		edges = append(edges, &Edge{Value: value})
	}
	return edges
}

func (fst *FST) nextID() int {
	id := fst.counter
	fst.counter = fst.counter + 1
	return id
}

func (fst *FST) newNode() Node {
	return Node{
		ID:    fst.nextID(),
		Edges: newEmptyEdges(),
	}
}

func (fst *FST) initialize() {
	start := fst.newNode()
	finish := fst.newNode()
	fst.Start = &start
	fst.Finish = &finish
}

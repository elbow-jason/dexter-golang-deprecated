package main

import (
	"errors"
	"fmt"
	// "log"
)

var notFoundErr = errors.New("Word Not Found")

var start Node = newNode()
var finish Node = newNode()

func constructFST(edges []*Edge) {
	var outterEdge *Edge
	var node *Node = &start
	for _, edge := range edges {
		if !node.hasValue(edge.Value) {
			node.addEdge(edge)
		} else {
			edge = node.getEdge(edge.Value)
		}
		node = edge.nextNode()
		outterEdge = edge
	}
	outterEdge.Next = &finish
}

func toEdges(word []byte) []*Edge {
	edges := make([]*Edge, 0)
	for _, value := range word {
		edges = append(edges, &Edge{Value: value})
	}
	return edges
}

func search(thing []byte) error {
	var edge *Edge
	var node *Node = &start
	for _, value := range thing {
		if node.hasValue(value) {
			edge = node.Edges[value]
			node = edge.nextNode()
		} else {
			return notFoundErr
		}
	}
	return nil
}

func findString(str string) {
	if err := search([]byte(str)); err != nil {
		fmt.Printf("%s: %q\n", err, str)
	} else {
		fmt.Printf("Word Found: %q\n", str)
	}
}

func main() {
	e1 := toEdges([]byte("mouse"))
	e2 := toEdges([]byte("mice"))

	constructFST(e1)
	constructFST(e2)
	findString("mouse")
	findString("mice")

	findString("hello")
}

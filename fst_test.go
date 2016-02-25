package dexter

import (
	"testing"
)

func TestNewFstIsInitialized(t *testing.T) {
	fst := NewFST()
	if fst.Start == nil {
		t.Fatal("A new FST should not have a nil Start.")
	}
	if fst.Finish == nil {
		t.Fatal("A new FST should not have a nil Finish.")
	}
	if fst.counter != 2 {
		t.Fatal("A new FST should have a counter that starts at 2 (after Start and Finish)")
	}
}

func TestFstNodeState(t *testing.T) {
	fst := NewFST()
	if fst.Start.ID != 0 {
		t.Fatalf("Expected Start node's ID to be 0. Got %d", fst.Start.ID)
	}
	if fst.Finish.ID != 1 {
		t.Fatalf("Expected Finish node's ID to be 1. Got %d", fst.Finish.ID)
	}
}

func TestFSTSearching(t *testing.T) {
	fst := NewFST()
	fst.AddString("mouse")

	if err := fst.Search("mouse"); err != nil {
		t.Fatal("Searching an FST for a term that is present should not return an error.")
	}
	if err := fst.Search("mice"); err == nil {
		t.Fatal("Searching an FST for a term that is not present should return an error.")
	}
}

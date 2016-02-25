package main

import (
	"fmt"
)

// Search for a string inside an FST
func Search(fst FST, str string) {
	if err := fst.Search(str); err != nil {
		fmt.Printf("%s: %q\n", err, str)
	} else {
		fmt.Printf("Word Found: %q\n", str)
	}
}

func main() {
	fst := newFST()
	fst.AddString("mouse")
	fst.AddString("mice")

	Search(fst, "mouse")
	Search(fst, "mice")
	Search(fst, "hello")
}

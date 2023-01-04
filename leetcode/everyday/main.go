package main

import (
	"fmt"

	"github.com/emirpasic/gods/trees/redblacktree"
)

func main() {
	tree := redblacktree.NewWithIntComparator()
	tree.Put(2, nil)
	tree.Put(1, nil)
	fmt.Print(tree.Keys())
}

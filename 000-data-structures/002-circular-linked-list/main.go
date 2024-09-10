package main

import "fmt"

func main() {
	aa := Node{Value: 1}
	bb := Node{Value: 2}
	cc := Node{Value: 3}

	aa.Next = &bb
	aa.Previous = &cc
	bb.Next = &cc
	bb.Previous = &aa
	cc.Next = &aa
	cc.Previous = &bb

	Print(&cc)
}

type Node struct {
	Value    int
	Next     *Node
	Previous *Node
}

func Print(root *Node) {
	printedNodes := make(map[*Node]bool)

	walk := root
	for walk != nil {
		if _, ok := printedNodes[walk]; ok {
			break
		}
		fmt.Println(walk.Value)
		printedNodes[walk] = true
		walk = walk.Next
	}
}

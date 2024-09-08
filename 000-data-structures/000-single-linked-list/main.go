package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	aa := Node{Value: 1}
	bb := Node{Value: 2}
	cc := Node{Value: 3}

	aa.Next = &bb
	bb.Next = &cc

	// -- Print --
	// Print(&aa)

	// -- Length --
	//n := Length(&aa)
	//fmt.Println(n)

	// -- Append --
	//Append(&aa, &Node{Value: 99})
	//Print(&aa)

	// -- Insert --
	//newNode, err := Insert(&aa, &Node{Value: 99}, 0)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	Print(newNode)
	//}

	// -- Delete --
	//err := Delete(&aa, 3)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	Print(&aa)
	//}
}

type Node struct {
	Value int
	Next  *Node
}

func Print(root *Node) {
	nodeWalk := root
	for nodeWalk.Next != nil {
		fmt.Println(nodeWalk.Value)
		nodeWalk = nodeWalk.Next
	}
	fmt.Println(nodeWalk.Value)
}

func Length(root *Node) int {
	nodeWalk := root
	counter := 1

	for nodeWalk.Next != nil {
		counter += 1
		nodeWalk = nodeWalk.Next
	}

	return counter
}

func Append(root *Node, newNode *Node) {
	nodeWalk := root
	for nodeWalk.Next != nil {
		nodeWalk = nodeWalk.Next
	}
	nodeWalk.Next = newNode
}

func Insert(root *Node, newNode *Node, loc int) (*Node, error) {
	nodeWalk := root
	counter := 0

	for nodeWalk.Next != nil {
		if counter == loc {
			temp := nodeWalk.Next
			newNode.Next = temp
			nodeWalk.Next = newNode
			return root, nil
		}
		counter += 1
		nodeWalk = nodeWalk.Next
	}

	if counter == loc {
		nodeWalk.Next = newNode
		return root, nil
	}

	return nil, errors.New("error insert")
}

func Delete(root *Node, loc int) error {
	nodeWalk := root
	previousWalk := root
	counter := 0

	for nodeWalk.Next != nil {
		if counter == loc {
			previousWalk.Next = nodeWalk.Next
			return nil
		}
		counter += 1
		previousWalk = nodeWalk
		nodeWalk = nodeWalk.Next
	}

	if counter == loc {
		previousWalk.Next = nil
	}

	return errors.New("error delete loc: " + strconv.Itoa(loc))
}

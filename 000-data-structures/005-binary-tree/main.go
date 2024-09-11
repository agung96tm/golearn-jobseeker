package main

import "fmt"

func main() {
	A := Node{data: "A"}
	B := Node{data: "B"}
	C := Node{data: "C"}
	D := Node{data: "D"}
	E := Node{data: "E"}
	F := Node{data: "F"}
	G := Node{data: "G"}

	A.left = &B
	A.right = &C
	B.left = &D
	B.right = &E
	C.left = &F
	C.right = &G

	//InOrder(&A)
	//PreOrder(&A)
	//PostOrder(&A)
}

type Node struct {
	data  string
	left  *Node
	right *Node
}

func InOrder(root *Node) {
	if root == nil {
		return
	}

	if root.left != nil {
		InOrder(root.left)
	}

	fmt.Println(root.data)

	if root.right != nil {
		InOrder(root.right)
	}
}

func PreOrder(root *Node) {
	if root == nil {
		return
	}

	fmt.Println(root.data)

	PreOrder(root.left)
	PreOrder(root.right)
}

func PostOrder(root *Node) {
	if root == nil {
		return
	}

	PostOrder(root.left)
	PostOrder(root.right)

	fmt.Println(root.data)
}

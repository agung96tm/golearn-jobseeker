package main

func main() {
	aa := Node{Value: 1}
	bb := Node{Value: 2}
	cc := Node{Value: 3}

	aa.Next = &bb
	bb.Next = &cc
	bb.Previous = &aa
	cc.Previous = &bb

	//newNode := Insert(&aa, &Node{Value: 99}, 3)
	//newNode := Delete(&aa, 2)
	//fmt.Println(newNode)
}

type Node struct {
	Value    int
	Next     *Node
	Previous *Node
}

func Insert(root *Node, newNode *Node, loc int) *Node {
	walk := root
	var prev *Node
	counter := 0

	if walk.Previous != nil {
		prev = walk.Previous
	}

	for walk != nil {
		if counter == loc {
			newNode.Next = walk
			newNode.Previous = prev

			if walk.Next != nil {
				walk.Next.Previous = newNode
			}

			if prev != nil {
				prev.Next = newNode
				return root
			}
		}

		prev = walk
		walk = walk.Next
		counter++
	}

	if counter == loc {
		newNode.Previous = prev
		prev.Next = newNode
		return root
	}

	return nil
}

func Delete(root *Node, loc int) *Node {
	counter := 0
	walk := root
	var prev *Node

	if walk.Previous != nil {
		prev = walk.Previous
	}

	for walk != nil {
		if counter == loc {
			if prev != nil {
				prev.Next = walk.Next
				if walk.Next != nil {
					prev.Next.Previous = prev
					return prev
				} else {
					return root
				}
			}

			if walk.Next != nil {
				temp := walk.Next
				walk.Next.Previous = nil
				return temp
			}
		}
		prev = walk
		walk = walk.Next
		counter++
	}

	return nil
}

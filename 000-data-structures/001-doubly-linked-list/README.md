Doubly Linked List
-------------------------

## Insert
```go
func main() {
    aa := Node{Value: 1}
    bb := Node{Value: 2}
    cc := Node{Value: 3}
    
    aa.Next = &bb
    bb.Next = &cc
    bb.Previous = &aa
    cc.Previous = &bb

    newNode := Insert(&aa, &Node{Value: 99}, 0)
    fmt.Println(newNode)
}
```

## Delete
```go
func main() {
    aa := Node{Value: 1}
    bb := Node{Value: 2}
    cc := Node{Value: 3}
    
    aa.Next = &bb
    bb.Next = &cc
    bb.Previous = &aa
    cc.Previous = &bb

    newNode := Delete(&aa, 0)
    fmt.Println(newNode)
}
```
Circular Linked List
-------------------------

## Print
```go
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

    Print(&aa)
}
```
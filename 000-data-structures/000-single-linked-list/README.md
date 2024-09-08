Single Linked List
-------------------------

## Print
```go
func main() {
    aa := Node{Value: 1}
    bb := Node{Value: 2}
    cc := Node{Value: 3}
    
    aa.Next = &bb
    bb.Next = &cc

    Print(&aa)
}
```

## Length
```go
func main() {
    aa := Node{Value: 1}
    bb := Node{Value: 2}
    cc := Node{Value: 3}
    
    aa.Next = &bb
    bb.Next = &cc

    Length(&aa)
}
```

## Append
```go
func main() {
    aa := Node{Value: 1}
    bb := Node{Value: 2}
    cc := Node{Value: 3}
    
    aa.Next = &bb
    bb.Next = &cc

    Append(&aa, &Node{Value: 99})
}
```

## Insert
```go
func main() {
    aa := Node{Value: 1}
    bb := Node{Value: 2}
    cc := Node{Value: 3}
    
    aa.Next = &bb
    bb.Next = &cc

    newNode, err := Insert(&aa, &Node{Value: 99}, 0)
    if err != nil {
    	fmt.Println(err)
    } else {
        // do something		
    }
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

    err := Delete(&aa, 3)
    if err != nil {
        fmt.Println(err)
    } else {
        // do something
    }
}
```
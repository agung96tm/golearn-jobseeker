Stack
-------------------------

## Push
```go
func main() {
    stack := NewStack()
    stack.Push(1)
    stack.Push(2)
    stack.Push(3)
	
	stack.Push(99) // 1, 2, 3
}
```

## Pop
```go
func main() {
    stack := NewStack()
    stack.Push(1)
    stack.Push(2)
    stack.Push(3)

    removedStack, err := stack.Pop(2)
    if err != nil {
	    fmt.Println(err)	
    } else {
	    fmt.Println(removedStack) // 3
	    fmt.Println(stack) // 1, 2
    }
}
```

## Insert
```go
func main() {
    stack := NewStack()
    stack.Push(1)
    stack.Push(2)
    stack.Push(3)
    
    err := stack.Insert(99, 1)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(stack) // 1, 99, 2, 3
    }
}
```
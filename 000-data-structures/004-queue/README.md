Queue
-------------------------

## Enqueue
```go
func main() {
    queue := NewQueue()
    queue.Enqueue(1)
    queue.Enqueue(2)
    queue.Enqueue(3)
	
    fmt.Println(queue) // 1, 2, 3
}
```

## Dequeue
```go
func main() {
    queue := NewQueue()
    queue.Enqueue(1)
    queue.Enqueue(2)
    queue.Enqueue(3)

    fmt.Println(queue) // 1, 2, 3

    first, err := queue.Dequeue()
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(first) // 1
        fmt.Println(queue) // 2, 3
    }	
}
```
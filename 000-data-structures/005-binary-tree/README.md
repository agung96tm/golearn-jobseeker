Binary Three
-------------------------

## InOrder
```go
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

    InOrder(&A) // D, B, E, A, F, C, G
}
```

## PreOrder
```go
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

    PreOrder(&A) // A, B, D, E, C, F, G
}
```


## PostOrder
```go
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

    PostOrder(&A) // D, E, B, F, G, C, A
}
```
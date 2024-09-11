Hashing
-------------------------

## GetHash
```go
func main() {
    aa := NewHashedMaps()
    aa.GetHash("mykey")
}
```


## Set
```go
func main() {
    aa := NewHashedMaps()
    aa.Set("mykey", "Sample Value")
}
```


## Get
```go
func main() {
    aa := NewHashedMaps()
    aa.Set("mykey", "Sample Value")
    
    fmt.Println(aa.Get("mykey")) // Sample Value
}
```
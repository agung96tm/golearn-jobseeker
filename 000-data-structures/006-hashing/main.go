package main

func main() {
	aa := NewHashedMaps()
	aa.Set("a", "Sample Value")
	//fmt.Println(aa.Get("a"))
}

type HashedMaps struct {
	items [100]string
}

func NewHashedMaps() HashedMaps {
	return HashedMaps{
		items: [100]string{},
	}
}

func (h *HashedMaps) GetHash(key string) int {
	totalSum := 0
	for _, v := range key {
		totalSum = totalSum + int(v)
	}

	hashKey := totalSum % len(h.items)
	return hashKey
}

func (h *HashedMaps) Set(key string, value string) {
	hashedKey := h.GetHash(key)
	h.items[hashedKey] = value
}

func (h *HashedMaps) Get(key string) string {
	hashedKey := h.GetHash(key)
	return h.items[hashedKey]
}

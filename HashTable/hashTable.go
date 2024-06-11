package main

import (
	"fmt"
)

type HashTable struct {
	size  int
	table []string
	keys  []string
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		size:  size,
		table: make([]string, size),
		keys:  make([]string, size),
	}
}

func (ht *HashTable) hash(key string) int {
	hash := 0
	for _, char := range key {
		hash = (hash*31 + int(char)) % ht.size
	}
	return hash
}

func (ht *HashTable) Insert(key, value string) {
	index := ht.hash(key)
	for ht.keys[index] != "" && ht.keys[index] != key {
		index = (index + 1) % ht.size
	}
	ht.table[index] = value
	ht.keys[index] = key
}

func (ht *HashTable) Retrieve(key string) string {
	index := ht.hash(key)
	for ht.keys[index] != "" {
		if ht.keys[index] == key {
			return ht.table[index]
		}
		index = (index + 1) % ht.size
	}
	return ""
}

func (ht *HashTable) Delete(key string) {
	index := ht.hash(key)
	for ht.keys[index] != "" {
		if ht.keys[index] == key {
			ht.table[index] = ""
			ht.keys[index] = ""
			return
		}
		index = (index + 1) % ht.size
	}
}

func main() {
	ht := NewHashTable(10)
	ht.Insert("name", "Alice")
	ht.Insert("age", "30")
	fmt.Println("Retrieve 'name':", ht.Retrieve("name"))
	ht.Delete("name")
	fmt.Println("Retrieve 'name' after deletion:", ht.Retrieve("name"))
}

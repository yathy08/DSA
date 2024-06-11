package main

import "fmt"

// Node represents a node in the doubly linked list
type Node struct {
    data int
    prev *Node
    next *Node
}

// DoublyLinkedList represents the doubly linked list
type DoublyLinkedList struct {
    head *Node
    tail *Node
}

// Insert adds a new node with the given data at the end of the list
func (dll *DoublyLinkedList) Insert(data int) {
    newNode := &Node{data: data}
    if dll.head == nil {
        dll.head = newNode
        dll.tail = newNode
    } else {
        dll.tail.next = newNode
        newNode.prev = dll.tail
        dll.tail = newNode
    }
}

// Traverse prints the elements of the list from head to tail
func (dll *DoublyLinkedList) Traverse() {
    current := dll.head
    for current != nil {
        fmt.Print(current.data, " -> ")
        current = current.next
    }
    fmt.Println("nil")
}

// ReverseTraverse prints the elements of the list from tail to head
func (dll *DoublyLinkedList) ReverseTraverse() {
    current := dll.tail
    for current != nil {
        fmt.Print(current.data, " -> ")
        current = current.prev
    }
    fmt.Println("nil")
}

func main() {
    dll := &DoublyLinkedList{}
    dll.Insert(1)
    dll.Insert(2)
    dll.Insert(3)
    dll.Insert(4)
    dll.Insert(5)

    fmt.Println("Traverse:")
    dll.Traverse() // Output: 1 -> 2 -> 3 -> 4 -> 5 -> nil

    fmt.Println("Reverse Traverse:")
    dll.ReverseTraverse() // Output: 5 -> 4 -> 3 -> 2 -> 1 -> nil
}

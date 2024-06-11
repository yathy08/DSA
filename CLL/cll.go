package main

import "fmt"

// Node represents a node in the circular linked list
type Node struct {
    data int
    next *Node
}

// CircularLinkedList represents the circular linked list
type CircularLinkedList struct {
    head *Node
}

// Insert adds a new node with the given data at the end of the list
func (cll *CircularLinkedList) Insert(data int) {
    newNode := &Node{data: data}
    if cll.head == nil {
        cll.head = newNode
        newNode.next = cll.head
    } else {
        current := cll.head
        for current.next != cll.head {
            current = current.next
        }
        current.next = newNode
        newNode.next = cll.head
    }
}

// Traverse prints the elements of the list
func (cll *CircularLinkedList) Traverse() {
    if cll.head == nil {
        fmt.Println("List is empty")
        return
    }
    current := cll.head
    for {
        fmt.Print(current.data, " -> ")
        current = current.next
        if current == cll.head {
            break
        }
    }
    fmt.Println("(back to head)")
}

func main() {
    cll := &CircularLinkedList{}
    cll.Insert(1)
    cll.Insert(2)
    cll.Insert(3)
    cll.Insert(4)
    cll.Insert(5)

    fmt.Println("Traverse:")
    cll.Traverse() // Output: 1 -> 2 -> 3 -> 4 -> 5 -> (back to head)
}

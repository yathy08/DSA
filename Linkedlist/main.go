package main 

import "fmt"

// Node represents an element in the linked list
type Node struct {
	data int
	next *Node
}

// LinkedList represents the linked list
type LinkedList struct {
	head *Node
}

// Insert adds a new node to the end of the list
func (ll *LinkedList) Insert(data int) {
	newNode := &Node{data: data}
	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

// Delete removes the first occurrence of the node with the given data
func (ll *LinkedList) Delete(data int) {
	if ll.head == nil {
		return
	}
	if ll.head.data == data {
		ll.head = ll.head.next
		return
	}
	current := ll.head
	for current.next != nil && current.next.data != data {
		current = current.next
	}
	if current.next != nil {
		current.next = current.next.next
	}
}

// Traverse prints all elements in the list
func (ll *LinkedList) Traverse() {
	current := ll.head
	for current != nil {
		fmt.Print(current.data, " -> ")
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	ll := &LinkedList{}
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Traverse() // Output: 1 -> 2 -> 3 -> nil

	ll.Delete(2)
	ll.Traverse() // Output: 1 -> 3 -> nil
}

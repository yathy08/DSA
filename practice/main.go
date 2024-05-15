package main

import "fmt"

type Node struct {
	data int
	next *Node
} 

type LinkedList struct {
	head *Node
}
func (ll *LinkedList) Insert(data int) {
	newnode := &Node {data: data}
	if ll.head == nil {
		ll.head = newnode
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		} 
		current.next = newnode
	}
}

func (ll *LinkedList) Traverse() {
	current := ll.head
	for current != nil {
		fmt.Print(current.data," -> ")
		current = current.next
	}
	fmt.Println(" nil ")
}

func main() {
	ll := &LinkedList{}
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Insert(4)
	ll.Insert(5)
	ll.Traverse()
}
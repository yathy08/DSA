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


func (ll *LinkedList) Delete(data int) {
	if ll.head == nil {
		return
	}
	if ll.head.data == data{
		ll.head = ll.head.next
		return
	}
	current := ll.head
	for current.next != nil {
		if current.next.data == data {
				current.next = current.next.next
				return
		}
		current = current.next
	}
}


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
	ll.Traverse() 

	ll.Delete(2)
	ll.Traverse() 
}

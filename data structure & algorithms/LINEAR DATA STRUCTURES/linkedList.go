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
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (ll *LinkedList) Display() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func (ll *LinkedList) Prepend(data int) {
	newNode := &Node{data: data, next: ll.head}
	ll.head = newNode
}

func (ll *LinkedList) InsertAfter(node *Node, data int) {
 
	if node == nil {
		return
	}
	
    newNode := &Node{data: data, next: node.next}
    node.next = newNode
 }

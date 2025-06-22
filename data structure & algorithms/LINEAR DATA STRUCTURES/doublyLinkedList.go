package main

import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (dll *DoublyLinkedList) Insert(data int) {
	newNode := &Node{data: data}
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
		return
	}
	newNode.prev = dll.tail
	dll.tail.next = newNode
	dll.tail = newNode
}
func (dll *DoublyLinkedList) Display() {
	current := dll.head
	for current != nil {
		fmt.Printf("%d <-> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func (dll *DoublyLinkedList) DisplayReverse() {
	current := dll.tail
	for current != nil {
		fmt.Printf("%d <-> ", current.data)
		current = current.prev
	}
	fmt.Println("nil")
}

func (dll *DoublyLinkedList) Prepend(data int) {
	newNode := &Node{data: data, next: dll.head}
	if dll.head != nil {
		dll.head.prev = newNode
	} else {
		dll.tail = newNode
	}
	dll.head = newNode
}
func (dll *DoublyLinkedList) Append(data int) {
	newNode := &Node{data: data, prev: dll.tail}
	if dll.tail != nil {
		dll.tail.next = newNode
	} else {
		dll.head = newNode
	}
	dll.tail = newNode
}

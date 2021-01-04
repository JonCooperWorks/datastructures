package datastructures

import (
	"fmt"
)

// LinkedList is a linked list as specified by https://www.tutorialspoint.com/data_structures_algorithms/linked_list_algorithms.htm
type LinkedList interface {
	Append(node *Node)
	Prepend(node *Node)
	Contains(value string) bool
	Delete(target *Node) error
	Traverse() []string
}

type linkedList struct {
	Head *Node
}

// NewLinkedList returns a linked list
func NewLinkedList() LinkedList {
	return &linkedList{
		Head: &Node{},
	}
}

// Append adds an item to the end of a linkedlist
func (l *linkedList) Append(node *Node) {
	current := l.Head
	for {
		if current.next == nil {
			current.next = node
			return
		}

		current = current.next
	}
}

// Prepend adds an element to the head of the linked list.
func (l *linkedList) Prepend(node *Node) {
	first := l.Head.next
	l.Head.next = node
	node.next = first
}

// Contains returns true if an element is found within a list.
func (l *linkedList) Contains(value string) bool {
	current := l.Head.next
	for {
		if current == nil {
			return false
		}

		if current.Value == value {
			return true
		}
		current = current.next
	}
}

func (l *linkedList) Delete(target *Node) error {
	current := l.Head
	for {
		if current.next == target {
			current.next = current.next.next
			return nil
		}

		if current.next == nil {
			return fmt.Errorf("node not in list")
		}

		current = current.next
	}

}

// Traverse returns a []string with all elements in a list.
func (l *linkedList) Traverse() []string {
	// Start at the node after the head.
	current := l.Head.next
	items := []string{}
	for {
		// If we hit the end of the list, return all collected items so far.
		if current == nil {
			return items
		}

		items = append(items, current.Value)
		current = current.next
	}
}

// A Node is an item in a LinkedList.
// It can contain any printable value.
type Node struct {
	Value string
	next  *Node
}

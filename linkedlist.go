package datastructures

import (
	"fmt"
	"sync"
)

// LinkedList is a linked list as specified by https://www.tutorialspoint.com/data_structures_algorithms/linked_list_algorithms.htm
type LinkedList interface {
	// Append adds an item to the end of a linkedlist
	Append(node *Node)

	// Prepend adds an element to the head of the linked list.
	Prepend(node *Node)

	// Contains returns true if an element is found within a list.
	Contains(value string) bool

	// Delete removes a node from the linked list.
	Delete(target *Node) error

	// Traverse returns a []string with all elements in a list.
	Traverse() []string
}

type linkedList struct {
	head *Node
	mu   sync.Mutex
}

// NewLinkedList returns a linked list
func NewLinkedList() LinkedList {
	return &linkedList{
		head: &Node{},
	}
}

func (l *linkedList) Append(node *Node) {
	l.mu.Lock()
	defer l.mu.Unlock()

	current := l.head
	for {
		if current.next == nil {
			current.next = node
			return
		}

		current = current.next
	}
}

func (l *linkedList) Prepend(node *Node) {
	l.mu.Lock()
	defer l.mu.Unlock()

	first := l.head.next
	l.head.next = node
	node.next = first
}

func (l *linkedList) Contains(value string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	current := l.head.next
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
	l.mu.Lock()
	defer l.mu.Unlock()

	current := l.head
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

func (l *linkedList) Traverse() []string {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Start at the node after the head.
	current := l.head.next
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

package datastructures

import (
	"fmt"
	"sync"
)

// DoubleLinkedList is a linked list where the nodes maintain a reference to both the previous and next node.
type DoubleLinkedList interface {
	// Append adds a node to the end of the list in O(1) time.
	Append(node *DoubleLinkedNode)

	// Prepend puts a new node at the beginning of the list in O(1) time
	Prepend(node *DoubleLinkedNode)

	// Contains returns true if a value is in a list, false otherwise in O(n) time.
	Contains(value string) bool

	// Delete removes an element from the list in O(n) time.
	Delete(target *DoubleLinkedNode) error

	// Traverse returns all elements in the list as a string slice in O(n) time.
	Traverse() []string

	// ReverseTraverse returns all elements in the list in reverse order as a string slice in O(n) time.
	ReverseTraverse() []string
}

// NewDoubleLinkedList returns a linked list that can be traversed in both directions.
func NewDoubleLinkedList() DoubleLinkedList {
	head := &DoubleLinkedNode{}
	return &doubleLinkedList{
		head: head,
		last: head,
	}
}

type doubleLinkedList struct {
	head *DoubleLinkedNode

	// maintain a pointer to the end of the list so we can ReverseTraverse faster and append in O(1) time.
	last *DoubleLinkedNode

	mu sync.Mutex
}

func (d *doubleLinkedList) Prepend(node *DoubleLinkedNode) {
	d.mu.Lock()
	defer d.mu.Unlock()

	next := d.head.next

	// Update last if we need to, otherwise leave it alone.
	if next == nil {
		d.last = node
	}
	node.next = next
	next.prev = node
	node.prev = d.head
	d.head.next = node
}

func (d *doubleLinkedList) Append(node *DoubleLinkedNode) {
	d.mu.Lock()
	defer d.mu.Unlock()

	// Do this in O(1) time using a pointer to the last node instead of traversing the list.
	node.prev = d.last
	d.last.next = node
	d.last = node
}

func (d *doubleLinkedList) Contains(value string) bool {
	d.mu.Lock()
	defer d.mu.Unlock()

	current := d.head.next
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

func (d *doubleLinkedList) Delete(target *DoubleLinkedNode) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	current := d.head
	for {
		if current == target {
			// Adjust the pointers around this node to remove the reference to it.
			current.prev.next = current.next
			current.next.prev = current.prev

			// Update the last pointer if we deleted the last node.
			if current.next == nil {
				d.last = current.prev
			}
			return nil
		}

		if current == nil {
			return fmt.Errorf("node not in list")
		}

		current = current.next
	}
}

func (d *doubleLinkedList) Traverse() []string {
	d.mu.Lock()
	defer d.mu.Unlock()

	current := d.head.next
	values := []string{}
	for {
		if current == nil {
			return values
		}

		values = append(values, current.Value)
		current = current.next
	}
}

func (d *doubleLinkedList) ReverseTraverse() []string {
	d.mu.Lock()
	defer d.mu.Unlock()

	// Start from the end of the list and go back to the start.
	current := d.last
	values := []string{}
	for {
		if current.prev == nil {
			return values
		}

		values = append(values, current.Value)
		current = current.prev
	}
}

// DoubleLinkedNode maintains a pointer to the previous and next nodes in the list.
type DoubleLinkedNode struct {
	prev  *DoubleLinkedNode
	next  *DoubleLinkedNode
	Value string
}

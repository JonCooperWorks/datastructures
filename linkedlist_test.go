package datastructures

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	expectedItems := []string{"pack", "test", "testing", "1245"}
	list := NewLinkedList()
	node1 := &Node{Value: "test"}
	node2 := &Node{Value: "testing"}
	node3 := &Node{Value: "1245"}
	node4 := &Node{Value: "pack"}
	list.Append(node1)
	list.Append(node2)
	list.Prepend(node4)
	list.Append(node3)

	if !reflect.DeepEqual(expectedItems, list.Traverse()) {
		t.Fatalf("expected %+v, got %+v", expectedItems, list.Traverse())
	}
}

func TestDelete(t *testing.T) {
	expectedItems := []string{"test", "testing", "1245"}
	list := NewLinkedList()
	node1 := &Node{Value: "test"}
	node2 := &Node{Value: "testing"}
	node3 := &Node{Value: "1245"}
	node4 := &Node{Value: "pack"}
	list.Append(node1)
	list.Append(node2)
	list.Prepend(node4)
	list.Append(node3)

	err := list.Delete(node4)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expectedItems, list.Traverse()) {
		t.Fatalf("expected %+v, got %+v", expectedItems[1:], list.Traverse())
	}
}

func TestContains(t *testing.T) {
	list := NewLinkedList()
	node1 := &Node{Value: "test"}
	node2 := &Node{Value: "testing"}
	node3 := &Node{Value: "1245"}
	node4 := &Node{Value: "pack"}
	list.Append(node1)
	list.Append(node2)
	list.Prepend(node4)
	list.Append(node3)

	if !list.Contains("test") {
		t.Fatalf("expected 'test' in list, %+v", list.Traverse())
	}

	err := list.Delete(node1)
	if err != nil {
		t.Fatal(err)
	}

	if list.Contains("test") {
		t.Fatalf("expected 'test' to be deleted from list, %+v", list.Traverse())
	}
}

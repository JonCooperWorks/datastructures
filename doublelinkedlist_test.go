package datastructures

import (
	"reflect"
	"testing"
)

func TestInsertDLL(t *testing.T) {
	expectedItems := []string{"pack", "test", "testing", "1245"}
	list := NewDoubleLinkedList()
	node1 := &DoubleLinkedNode{Value: "test"}
	node2 := &DoubleLinkedNode{Value: "testing"}
	node3 := &DoubleLinkedNode{Value: "1245"}
	node4 := &DoubleLinkedNode{Value: "pack"}
	list.Append(node1)
	list.Append(node2)
	list.Prepend(node4)
	list.Append(node3)

	if !reflect.DeepEqual(expectedItems, list.Traverse()) {
		t.Fatalf("expected %+v, got %+v", expectedItems, list.Traverse())
	}

}

func TestReverseTraverse(t *testing.T) {
	expectedItems := []string{"1245", "testing", "test", "pack"}
	list := NewDoubleLinkedList()
	node1 := &DoubleLinkedNode{Value: "test"}
	node2 := &DoubleLinkedNode{Value: "testing"}
	node3 := &DoubleLinkedNode{Value: "1245"}
	node4 := &DoubleLinkedNode{Value: "pack"}
	list.Append(node1)
	list.Append(node2)
	list.Prepend(node4)
	list.Append(node3)

	if !reflect.DeepEqual(expectedItems, list.ReverseTraverse()) {
		t.Fatalf("expected %+v, got %+v", expectedItems, list.ReverseTraverse())
	}

}

func TestDeleteDLL(t *testing.T) {
	expectedItems := []string{"test", "testing", "1245"}
	list := NewDoubleLinkedList()
	node1 := &DoubleLinkedNode{Value: "test"}
	node2 := &DoubleLinkedNode{Value: "testing"}
	node3 := &DoubleLinkedNode{Value: "1245"}
	node4 := &DoubleLinkedNode{Value: "pack"}
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

func TestContainsDLL(t *testing.T) {
	list := NewDoubleLinkedList()
	node1 := &DoubleLinkedNode{Value: "test"}
	node2 := &DoubleLinkedNode{Value: "testing"}
	node3 := &DoubleLinkedNode{Value: "1245"}
	node4 := &DoubleLinkedNode{Value: "pack"}
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

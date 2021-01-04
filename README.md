# Data Structures
This repository has implementations of data structures in Go.
They are purely for my own learning, do not use these in real software.

### Linked List
A linked list is data structure that stores data as a set of nodes.
Each node contains a link to the node after it, called `next`, until a `nil` `next` node signifies the end of the list.
The linked list in this package can only hold strings.
Linked lists allow the following operations:

+ `Append`
  
  `Append` adds an node to the end of the Linked List.
  It is a `O(n)` operation, since the list has to be fully traversed to add a new node at the end.

+ `Prepend`
  
  `Prepend` adds a node to the beginning of the linked list.
  It runs in `O(1)` time, since it simply modifies the `next` pointer at the head of the list.

+ `Contains`
  
  `Contains` returns `true` if a Linked List contains a particular value, or `false` if it doesn't.
  Since `Contains` requires traversing the list, it is a `O(n)` operation.

+ `Delete`
  
  `Delete` removes a node from a list by adjusting the `next` pointer of the node before it.
  This is a `O(n)` operation, since the list has to be traversed to find the target.

+ `Traverse`
  
  `Traverse` returns the values in a Linked List as a `[]string` slice.
  `Traverse` is done in `O(n)` time, as every node in the list must be visited.
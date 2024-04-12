package main

import (
	"fmt"
)

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) Insert(value int) *LinkedList {
	if ll.head == nil {
		ll.head = NewNode(value)
	} else {
		newNode := NewNode(value)
		current := ll.head

		for current.next != nil {
			current = current.next
		}

		current.next = newNode
		newNode.prev = current
	}
	return ll
}

func (ll *LinkedList) Traverse() {
	fmt.Print("[")
	for ll.head != nil {
		fmt.Printf(" %d ", ll.head.value)
		ll.head = ll.head.next
	}
	fmt.Print("]\n")
}

/*
	Given a Singly Linked List, the task is to find the middle of the linked list. If the number of nodes are even, then there would be two middle nodes, so return the second middle node.

	Example:

	Input: linked list = 1 -> 2 -> 3 -> 4 -> 5
	Output: 3
	Explanation: There are 5 nodes in the linked list and there is one middle node whose value is 3.

	Input: linked list = 1 -> 2 -> 3 -> 4 -> 5 -> 6
	Output: 4
	Explanation: There are 6 nodes in the linked list, so we have two middle nodes: 3 and 4, but we will return the second middle node which is 4.

	@Source: geekforgeeks

*/

func (ll *LinkedList) Middle() {
	var list []int

	for ll.head != nil {
		list = append(list, ll.head.value)
		ll.head = ll.head.next
	}

	var middle int
	listDivBy2 := (len(list) / 2)

	if len(list) > 0 {
		middle = list[listDivBy2]
	}

	fmt.Println("Middle =", middle)
}

type Node struct {
	value int
	next  *Node
	prev  *Node
}

func NewNode(value int) *Node {
	return &Node{
		value: value,
		next:  nil,
		prev:  nil,
	}
}

func main() {
	ll := &LinkedList{}
	ll.Insert(1).
		Insert(2).
		Insert(3).
		Insert(4).
		Insert(5)

	// ll.Traverse()

	ll.Middle()
}

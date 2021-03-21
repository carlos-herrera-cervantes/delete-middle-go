package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	length int
	head   *node
	tail   *node
}

// Returns the length of linked list
func (list linkedList) getLength() int {
	return list.length
}

// Prints the elements of linked list
func (list linkedList) printList() {
	for list.head != nil {
		fmt.Printf("%v -> ", list.head.data)
		list.head = list.head.next
	}

	fmt.Println()
}

// Adds new node in the linked list
func (list *linkedList) pushBack(node *node) {
	if list.head == nil {
		list.head = node
		list.tail = node
		list.length++

		return
	}

	list.tail.next = node
	list.tail = node
	list.length++
}

// Deletes the middle node
func (list *linkedList) deleteMiddle() {
	head := list.head

	if head == nil {
		return
	}

	if head.next == nil {
		return
	}

	middle := list.length / 2

	for middle > 1 {
		head = head.next
		middle--
	}

	head.next = head.next.next
}

func main() {
	node1 := &node{data: 10}
	node2 := &node{data: 3}
	node3 := &node{data: 11}
	node4 := &node{data: 100}
	linkedList1 := linkedList{}

	linkedList1.pushBack(node1)
	linkedList1.pushBack(node2)
	linkedList1.pushBack(node3)
	linkedList1.pushBack(node4)
	linkedList1.deleteMiddle()

	linkedList1.printList()
}

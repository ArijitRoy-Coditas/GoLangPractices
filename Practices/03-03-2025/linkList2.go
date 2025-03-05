package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkList struct {
	head *node
}

func (list *linkList) insertAtBeginning(data int) {
	newNode := &node{data: data}
	newNode.next = list.head
	list.head = newNode
}

func add(l1, l2 *linkList) *linkList {
	res := &node{}
	current := res
	t1 := l1.head
	t2 := l2.head
	carry := 0
	for t1 != nil || t2 != nil || carry > 0 {
		sum := carry

		if t1 != nil {
			sum = sum + t1.data
			t1 = t1.next
		}

		if t2 != nil {
			sum = sum + t2.data
			t2 = t2.next
		}

		newNode := &node{data: sum % 10}
		carry = sum / 10

		current.next = newNode
		current = current.next

	}
	if carry != 0 {
		newNode := &node{data: carry}
		current.next = newNode
	}
	return &linkList{head: res.next}

}

func (list *linkList) display() {
	current := list.head

	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}
	fmt.Println()
}

func main() {
	myList := linkList{}
	myList.insertAtBeginning(9)
	myList.insertAtBeginning(9)
	myList.insertAtBeginning(9)
	myList.insertAtBeginning(9)

	myList.insertAtBeginning(9)
	myList.insertAtBeginning(9)
	myList.insertAtBeginning(9)

	list1 := linkList{}
	list1.insertAtBeginning(9)
	list1.insertAtBeginning(9)
	list1.insertAtBeginning(9)
	list1.insertAtBeginning(9)

	result := add(&myList, &list1)
	result.display()

}

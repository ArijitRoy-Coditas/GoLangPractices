package main

import "fmt"

type node struct {
	data int
	next *node
}

type LinkList struct {
	head *node
}

func (list *LinkList) insertAtBeginning(data int) {
	newNode := &node{data: data}
	newNode.next = list.head
	list.head = newNode
}

func (list *LinkList) insertAtEnd(data int) {
	newNode := &node{data: data}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head

	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func (list *LinkList) display() {
	current := list.head

	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}
}

// func (list *LinkList) display(node *node) {
// 	if node == nil {
// 		return
// 	}
// 	list.display(node.next)
// 	fmt.Print(node.data, " ")
// }

// func (list *LinkList) displayRev() {
// 	list.display(list.head)
// }

// Display in reverse
func (list *LinkList) sumList(node *node) int {
	if node == nil {
		return 0
	}
	var sum int
	sum = list.sumList(node.next)
	return sum*10 + node.data
}

func (list *LinkList) sumRev() int {
	sum := list.sumList(list.head)
	return sum
}

func main() {
	myList := LinkList{}
	myList.insertAtBeginning(9)
	myList.insertAtBeginning(9)
	myList.insertAtBeginning(9)
	myList.insertAtBeginning(9)

	myList.insertAtBeginning(9)
	myList.insertAtBeginning(9)
	myList.insertAtBeginning(9)

	sum1 := myList.sumRev()

	list1 := LinkList{}
	list1.insertAtBeginning(9)
	list1.insertAtBeginning(9)
	list1.insertAtBeginning(9)
	list1.insertAtBeginning(9)

	sum2 := list1.sumRev()
	newList := LinkList{}
	total := sum1 + sum2
	sum := 0
	for total > 0 {
		sum = total % 10
		fmt.Println(sum)
		newList.insertAtEnd(sum)
		total = total / 10
	}
	newList.display()
}

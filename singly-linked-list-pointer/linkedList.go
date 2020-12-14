package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// Node represents a node of linked list
type Node struct {
	value int
	next  *Node
}

// LinkedList represents a linked list
type LinkedList struct {
	head *Node
	len  int
}

// Init empty LinkedList
func newLinkedList() *LinkedList {
	return &LinkedList{}
}

// add to the head
func (l *LinkedList) insertAtHead(val int) {
	temp := new(Node)
	temp.value = val
	temp.next = l.head

	l.head = temp
	l.len++
	return
}

// insert new node at the end o linked list
// For adding a new node at the end of linked list
// we have to traverse through the linked list till end.
func (l *LinkedList) insertAtEnd(val int) {
	node := new(Node)
	node.value = val
	if l == nil {
		l.head = node
		l.len++
		return
	}
	ptr := l.head
	// traverse through the linked list
	// until the end then add new one.
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = node
			l.len++
			return
		}
		ptr = ptr.next
	}
}

// insert a node to the specific position
func (l *LinkedList) insertAtPos(pos, val int) error {
	node := new(Node)
	node.value = val

	if pos < 0 || pos > l.len {
		fmt.Println("Pos can'nt be negative or bigger than length")
		return errors.New("Pos can'nt be negative or bigger than length")
	}
	if pos == 0 {
		l.insertAtHead(val)
		return nil
	}

	ptrPrev, _ := l.getAtPos(pos - 1)
	ptrNex, _ := l.getAtPos(pos)

	ptrPrev.next = node
	node.next = ptrNex
	l.len++
	return nil
}

// del a node at a specific position
func (l *LinkedList) deleteAtPos(pos int) error {
	if l.len == 0 {
		fmt.Println("Linkedlist is empty")
		return errors.New("Linkedlist is empty")
	}
	if pos < 0 || pos > l.len {
		fmt.Println("Pos can'nt be negative or bigger than length")
		return errors.New("Pos can'nt be negative or bigger than length")
	}
	if pos == 0 {
		l.head, _ = l.getAtPos(pos + 1)
		l.len--
		return nil
	}

	ptrPrev, _ := l.getAtPos(pos - 1)
	ptrNex, _ := l.getAtPos(pos + 1)

	ptrPrev.next = ptrNex
	l.len--
	return nil
}

// delete the nodes with the specific val in the Linkedlist
func (l *LinkedList) deleteVal(val int) error {
	if l.len == 0 {
		fmt.Println("Linkedlist is empty")
		return errors.New("Linkedlist is empty")
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.value == val {
			if i > 0 {
				ptrPrev, _ := l.getAtPos(i - 1)
				ptrPrev.next = ptr.next
			}
			l.head = l.head.next
		}
		ptr = ptr.next
	}
	return nil
}

// get node at the specfic position
func (l *LinkedList) getAtPos(pos int) (*Node, error) {
	ptr := l.head
	// if pos < 0 return head of LinkedList
	if pos < 0 {
		fmt.Println("Position can not be negative")
		return nil, errors.New("Position can not be negative")
	}
	if pos > l.len-1 {
		fmt.Println("Position is bigger than length")
		return nil, errors.New("Position is bigger than length")
	}

	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}
	return ptr, nil
}

func (l *LinkedList) display() {
	var temp *Node
	temp = l.head
	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}

}
func main() {
	// var l *LinkedList
	l := newLinkedList()
	fmt.Println("=== insert 5 data===")
	n := 0
	for n < 5 {
		fmt.Printf("count %d\n", n)
		value := rand.Intn(100)
		l.insertAtHead(value)
		fmt.Printf("data %d\n", value)
		n++
	}
	fmt.Println("== display ====")
	fmt.Println(l)
	l.display()
	fmt.Println("------------------")
	l.insertAtPos(3, 9999)
	l.display()
	fmt.Println("------------------")
	l.deleteAtPos(0)
	l.display()
	fmt.Println("------------------")
	l.getAtPos(1000)
	l.display()
	fmt.Println("------------------")
	l.deleteVal(9999)
	l.display()
}

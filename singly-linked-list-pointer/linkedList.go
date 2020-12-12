package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	value int
	next  *Node
}

// add to the head
func add(list *Node, data int) *Node {
	if list == nil {
		list := new(Node)
		list.value = data
		list.next = nil

		return list
	}
	temp := new(Node)
	temp.value = data
	temp.next = list

	list = temp

	return list
}

func display(list *Node) {
	var temp *Node
	temp = list
	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}
}
func main() {
	var head *Node
	head = nil
	fmt.Println("=== insert 5 data===")
	n := 0
	for n < 5 {
		fmt.Printf("data %d\n", n)
		value := rand.Intn(100)
		head = add(head, value)
		fmt.Printf("data %d\n", value)
		n++
	}
	fmt.Println("== display ====")
	display(head)
}

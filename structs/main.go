package main

import "fmt"

func main() {
	type Node struct {
		value int
		next  *Node
	}
	
	type LinkedList struct {
		head *Node
	}
	
	func (l *LinkedList) Insert(value int) {
		newNode := &Node{value: value}
		if l.head == nil {
			l.head = newNode
		} else {
			current := l.head
			for current.next != nil {
				current = current.next
			}
			current.next = newNode
		}
	}
	
}

// to execute us below command
// go run main.go

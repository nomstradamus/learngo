package main

import (
	"fmt"
)

type LinkList struct {
	head *Node
	tail *Node
}

type Node struct {
	i    int
	next *Node
}

type DataProvider interface {
	getData() DataProvider
	putData(DataProvider)
}

func (l *LinkList) addNode(data int) {

	n := Node{i: data, next: nil}

	if l.head == nil {
		l.head = &n
		l.tail = &n
	}
	l.tail.next = &n
	l.tail = &n

}

func main() {
	l := LinkList{
		head: nil,
		tail: nil,
	}

	l.addNode(2)
	l.addNode(3)

	for x := l.head; x != nil; x = x.next {
		fmt.Printf("%d -->", x.i)
	}

}

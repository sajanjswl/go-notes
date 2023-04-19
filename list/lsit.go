package main

import "fmt"

type Node struct {
	RightNode *Node
	Key       interface{}
	LeftNode  *Node
}

type List struct {
	Head *Node
	Tail *Node
}

func (l *List) Add(key interface{}) {

	node := &Node{
		Key: key,
	}

	if l.Head == nil || l.Tail == nil {
		l.Head = node
		l.Tail = l.Head

		fmt.Println("Printing l.Head", l.Head.Key)
		fmt.Println("Printing l.Tail", l.Tail.Key)

	} else {
		l.Head.RightNode = node

		l.Head.LeftNode =
			fmt.Println("fmt.Pritning head", l.Head.Key)
		fmt.Println("fmt.Pritning Tail", l.Tail.Key)
		fmt.Println("fmt.Pritning head.RightNode", l.Head.RightNode.Key)

	}

}

func (l *List) AddDobuleList(key interface{}) {

	node := &Node{
		Key: key,
	}

	if l.Head == nil || l.Tail == nil {
		l.Head = node
		l.Tail = l.Head

	} else {

		node.LeftNode = l.Head
		l.Head.RightNode = node
		l.Head = l.Head.RightNode

	}

}

func (l *List) Show() {

	node := l.Tail

	for node != nil {
		fmt.Println("key -> %v", node.Key)
		node = node.RightNode
	}

}
func main() {

	list := &List{}

	for i := 1; i <= 10; i++ {
		list.AddDobuleList(i)
		// list.Add(i)
	}

	list.Show()

}

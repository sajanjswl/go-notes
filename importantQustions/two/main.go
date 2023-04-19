package main

import (
	"fmt"
)

// Single Linked List
type Node struct {
	Key      int
	NextNode *Node
}

func NewNode(key int) *Node {

	var node Node
	node.Key = key
	node.NextNode = nil
	return &node
}

func Show(head *Node) {
	temHead := head

	for temHead != nil {
		fmt.Println(temHead.Key)
		temHead = temHead.NextNode
	}

}


// BST Create
type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int64
}

type BinaryTree struct {
	root *BinaryNode
}

func (t *BinaryTree) insert(data int64) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryNode{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *BinaryNode) insert(data int64) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}

func main() {
	head := NewNode(1)
	tempHead := head
	for i := 2; i < 10; i++ {
		if tempHead.NextNode == nil {
			tempHead.NextNode = NewNode(i)
		}
		tempHead = tempHead.NextNode

	}

	Show(head)
}









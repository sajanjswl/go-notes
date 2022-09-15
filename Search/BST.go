package main

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) InOrderTraverse(array []int) []int {

	if tree == nil {
		return array
	}
	array = tree.Left.InOrderTraverse(array)
	array = append(array, tree.Value)
	array = tree.Right.InOrderTraverse(array)

	// Write your code here.
	return array
}

func (tree *BST) PreOrderTraverse(array []int) []int {
	if tree == nil {
		return array
	}

	array = append(array, tree.Value)
	array = tree.Left.PreOrderTraverse(array)
	array = tree.Right.PreOrderTraverse(array)

	// Write your code here.
	return array
}

func (tree *BST) PostOrderTraverse(array []int) []int {
	if tree == nil {
		return array
	}

	array = tree.Left.PostOrderTraverse(array)
	array = tree.Right.PostOrderTraverse(array)
	array = append(array, tree.Value)
	// Write your code here.
	return array
}

package main

import "fmt"

func main() {
	fmt.Println("Define the tree")

	arr := []int{2, 3, 4, 7, 5, 8, 9, 8, 10, 12, 11, 13, 8}

	l := len(arr)
	root := &node{}

	s := make([]int, 0)

	tree := constructBinTree(arr, root, 0, l)
	result := preOrderTraverse(tree, s)

	fmt.Println(result)
}

type node struct {
	key   int
	left  *node
	right *node
}

// let's construct a simple Binary tree
func constructBinTree(arr []int, root *node, index int, l int) *node {

	var leftIndex int
	var rigntIndex int

	// base condition
	if index < l {
		n := &node{key: arr[index]}
		root = n

		leftIndex = 2*index + 1
		rigntIndex = 2*index + 2

		root.left = constructBinTree(arr, root.left, leftIndex, l)
		root.right = constructBinTree(arr, root.right, rigntIndex, l)
	}
	return root
}

func preOrderTraverse(tree *node, array []int) []int {
	if tree != nil {
		array = append(array, tree.key)
		preOrderTraverse(tree.left, array)
		preOrderTraverse(tree.right, array)
	}

	return array
}

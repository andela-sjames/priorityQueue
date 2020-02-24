package main

import "fmt"

func main() {
	fmt.Println("Define the array")

	arrOne := []int{2, 3, 4, 7, 5, 8, 9, 8, 10, 12, 11, 13, 8} // Binary Array
	arrTwo := []int{4, 7, 8, 3, 2, 6, 5}                       // Non Heap Array

	valTwo := maxHeapify(arrTwo, 0, len(arrTwo))
	fmt.Println(valTwo)

	fmt.Println(arrOne)

	l := len(arrOne)
	root := &node{}

	s := make([]int, 0)

	fmt.Println("Construct the Binary tree")
	tree := constructBinTree(arrOne, root, 0, l)

	fmt.Println("preOrderTraverse the Binary tree")
	result := preOrderTraverse(tree, s)

	fmt.Println(result)
}

type node struct {
	key   int
	left  *node
	right *node
}

func maxHeapify(arr []int, rootIndex, arrayLength int) []int {
	var leftIndex int
	var rightIndex int

	var largest int

	leftIndex = 2*rootIndex + 1
	rightIndex = 2*rootIndex + 2

	if leftIndex <= arrayLength && arr[leftIndex] > arr[rootIndex] {
		largest = leftIndex
	} else {
		largest = rootIndex
	}

	if rightIndex <= arrayLength && arr[rightIndex] > arr[largest] {
		largest = rightIndex
	}

	if largest != rootIndex {
		swap(arr, rootIndex, largest)
		maxHeapify(arr, largest, arrayLength)
	}

	return arr
}

func swap(arr []int, x, y int) {
	// swap here dude.
	tmp := arr[x]
	arr[x] = arr[y]
	arr[y] = tmp

}

// let's construct a simple Binary tree
func constructBinTree(arr []int, root *node, index int, l int) *node {

	var leftIndex int
	var rightIndex int

	// base condition
	if index < l {
		n := &node{key: arr[index]}
		root = n

		leftIndex = 2*index + 1
		rightIndex = 2*index + 2

		root.left = constructBinTree(arr, root.left, leftIndex, l)
		root.right = constructBinTree(arr, root.right, rightIndex, l)
	}
	return root
}

func preOrderTraverse(tree *node, array []int) []int {

	if tree != nil {
		array = append(array, tree.key)             // root
		array = preOrderTraverse(tree.left, array)  // left
		array = preOrderTraverse(tree.right, array) // right
	}

	return array
}

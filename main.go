package main

import "fmt"

func main() {
	fmt.Println("Define the array")

	arrOne := []int{2, 3, 4, 7, 5, 8, 9, 8, 10, 12, 11, 13, 8} // Binary Array
	// arrTwo := []int{4, 7, 8, 3, 2, 6, 5}                       // Non Heap Array

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

// To heapify a subtree rooted with node i (rootIndex)
// which is an index in arr[]. arrayLength is the size of the heap.
func heapify(arr []int, rootIndex, arrayLength int) {
	var leftIndex int
	var rightIndex int

	var largest int

	largest = rootIndex

	leftIndex = 2*rootIndex + 1
	rightIndex = 2*rootIndex + 2

	// if the left child is larger than root
	if leftIndex <= arrayLength && arr[leftIndex] > arr[rootIndex] {
		largest = leftIndex
	}

	// if the right child is larger than largest so far
	if rightIndex <= arrayLength && arr[rightIndex] > arr[largest] {
		largest = rightIndex
	}

	// if the largest is not root
	if largest != rootIndex {
		swap(arr, rootIndex, largest)

		// recursively heapify the affected sub-tree
		heapify(arr, largest, arrayLength)
	}
}

func buildHeap(arr []int, size int) {

	// Index of the last non-leaf node
	startIdx := (size / 2) - 1

	// Perform reverse level order traversal
	// from last non-leaf node and heapify
	// each node
	for i := startIdx; i >= 0; i-- {
		heapify(arr, i, size)
	}

}

func swap(arr []int, x, y int) {
	// swap here dude.
	tmp := arr[x]
	arr[x] = arr[y]
	arr[y] = tmp
}

// construct a simple Binary tree
func constructBinTree(arr []int, root *node, index int, l int) *node {

	var leftIndex int
	var rightIndex int

	// base condition
	if index < l {
		n := &node{key: arr[index]}
		root = n

		leftIndex = 2*index + 1
		rightIndex = 2*index + 2

		// recursive steps
		root.left = constructBinTree(arr, root.left, leftIndex, l)
		root.right = constructBinTree(arr, root.right, rightIndex, l)
	}
	return root
}

// Depth first search (DFS) Pre-Order traversal
func preOrderTraverse(tree *node, array []int) []int {

	if tree != nil {
		array = append(array, tree.key)             // root
		array = preOrderTraverse(tree.left, array)  // left
		array = preOrderTraverse(tree.right, array) // right
	}

	return array
}

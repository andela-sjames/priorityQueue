package main

import "fmt"

func main() {
	fmt.Println("Define the array")

	heapArr := []int{4, 7, 8, 3, 2, 6, 5, 1, 3, 5, 4, 6, 13, 10, 9, 8, 15, 17} // Non Heap Array

	buildHeap(heapArr, len(heapArr))
	printHeap(heapArr, len(heapArr))
}

type node struct {
	key   int
	left  *node
	right *node
}

// To heapify a subtree rooted with node i (rootIndex)
// which is an index in arr[]. size is the size of the heap.
func heapify(arr []int, rootIndex, size int) {
	var leftIndex int
	var rightIndex int

	var largest int

	largest = rootIndex

	leftIndex = 2*rootIndex + 1
	rightIndex = 2*rootIndex + 2

	// if the left child is larger than root
	if leftIndex < size && arr[leftIndex] > arr[largest] {
		largest = leftIndex
	}

	// if the right child is larger than largest so far
	if rightIndex < size && arr[rightIndex] > arr[largest] {
		largest = rightIndex
	}

	// if the largest is not root
	if largest != rootIndex {
		swap(arr, rootIndex, largest)

		// recursively heapify the affected sub-tree
		heapify(arr, largest, size)
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

func printHeap(arr []int, size int) {
	result := make([]int, 0)
	for _, val := range arr {
		result = append(result, val)
	}

	fmt.Println(result)
}

func swap(arr []int, x, y int) {
	// swap here dude.
	tmp := arr[x]
	arr[x] = arr[y]
	arr[y] = tmp
}

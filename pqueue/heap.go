package pqueue

import (
	"encoding/json"
	"fmt"
)

// MaxHeap struct defined to describe the Priority Queue ADT
type MaxHeap struct {
	pqArr []*pItem
	table map[int][]int
}

// NewHeap returns a new MaxHeap struct
func NewHeap() *MaxHeap {
	return &MaxHeap{
		table: make(map[int][]int),
	}
}

type node struct {
	key   int
	left  *node
	right *node
}

// pItem struct defined to describe a priority object
type pItem struct {
	Item     string
	Priority int
}

// InsertPriority inserts a new item with priority in integer
func (m *MaxHeap) InsertPriority(item string, priority int) {
	// Insert an Item at the end of array and bubble back up to
	// satisfy the heap invariant.

	newPriority := &pItem{Item: item, Priority: priority}
	m.pqArr = append(m.pqArr, newPriority)
	m.buildHeap(m.pqArr, len(m.pqArr))
}

// ShowPriority returns the highest priority but does
// not remove it from the priority queue
func (m *MaxHeap) ShowPriority() (string, int) {
	priorityOne := m.pqArr[0]

	// fmt.Printf("%v\n", priorityOne)
	return priorityOne.Item, priorityOne.Priority
}

// buildHeap function defined
func (m *MaxHeap) buildHeap(arr []*pItem, size int) {

	// Index of the last non-leaf node
	startIdx := (size / 2) - 1

	// Perform reverse level order traversal
	// from last non-leaf node and heapify
	// each node
	for i := startIdx; i >= 0; i-- {
		m.heapify(arr, i, size)
	}
}

func (m *MaxHeap) hashHeapify(arr []*pItem, rootIndex, size int) {
	var leftIndex int
	var rightIndex int

	var largest int

	largest = rootIndex

	leftIndex = 2*rootIndex + 1
	rightIndex = 2*rootIndex + 2

	// if the left child is larger than root
	if leftIndex < size && arr[leftIndex].Priority > arr[largest].Priority {
		largest = leftIndex
	}

	// if the right child is larger than largest so far
	if rightIndex < size && arr[rightIndex].Priority > arr[largest].Priority {
		largest = rightIndex
	}

	// if the largest is not root
	if largest != rootIndex {
		m.swap(arr, rootIndex, largest)

		// recursively heapify the affected sub-tree
		m.heapify(arr, largest, size)
	}
}

// To heapify a subtree rooted with node i (rootIndex)
// which is an index in arr[]. size is the size of the heap.
func (m *MaxHeap) heapify(arr []*pItem, rootIndex, size int) {
	var leftIndex int
	var rightIndex int

	var largest int

	largest = rootIndex

	leftIndex = 2*rootIndex + 1
	rightIndex = 2*rootIndex + 2

	// if the left child is larger than root
	if leftIndex < size && arr[leftIndex].Priority > arr[largest].Priority {
		largest = leftIndex
	}

	// if the right child is larger than largest so far
	if rightIndex < size && arr[rightIndex].Priority > arr[largest].Priority {
		largest = rightIndex
	}

	// if the largest is not root
	if largest != rootIndex {
		m.swap(arr, rootIndex, largest)

		// recursively heapify the affected sub-tree
		m.heapify(arr, largest, size)
	}
}

// PrintHeap function defined
func (m *MaxHeap) PrintHeap() {
	result := make([]string, 0)

	for _, val := range m.pqArr {
		vjson, _ := json.Marshal(val)
		// fmt.Println(string(vjson))
		result = append(result, string(vjson))
	}
	fmt.Println(result)
}

func (m *MaxHeap) swap(arr []*pItem, x, y int) {
	// swap here dude.
	tmp := arr[x]
	arr[x] = arr[y]
	arr[y] = tmp
}

// RemovePriority defined to remove an item from the heap
func RemovePriority() {

}

// PollPriority defined to get the top element from the heap
func PollPriority() {

}

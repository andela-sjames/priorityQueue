package pqueue

import "fmt"

// MaxHeap struct defined to describe the Priority Queue ADT
type MaxHeap struct {
	pqArr []*PItem
}

type node struct {
	key   int
	left  *node
	right *node
}

// PItem struct defined to describe a priority object
type PItem struct {
	Item     string
	Priority int
}

// InsertPriority insert a new item with priority in integer
func (m *MaxHeap) InsertPriority(item string, priority int) {
	// Insert an Item at the end of array and bubble back up to
	// satisfy the heap invariant.

	newPriority := &PItem{Item: "Visit China", Priority: 4}
	m.pqArr = append(m.pqArr, newPriority)

	m.BuildHeap(m.pqArr, len(m.pqArr))

}

// To heapify a subtree rooted with node i (rootIndex)
// which is an index in arr[]. size is the size of the heap.
func (m *MaxHeap) heapify(arr []*PItem, rootIndex, size int) {
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

// BuildHeap function defined
func (m *MaxHeap) BuildHeap(arr []*PItem, size int) {

	// Index of the last non-leaf node
	startIdx := (size / 2) - 1

	// Perform reverse level order traversal
	// from last non-leaf node and heapify
	// each node
	for i := startIdx; i >= 0; i-- {
		m.heapify(arr, i, size)
	}

}

// PrintHeap function defined
func (m *MaxHeap) PrintHeap(arr []PItem, size int) {
	result := make([]PItem, 0)
	for _, val := range arr {
		result = append(result, val)
	}

	fmt.Println(result)
}

func (m *MaxHeap) swap(arr []*PItem, x, y int) {
	// swap here dude.
	tmp := arr[x]
	arr[x] = arr[y]
	arr[y] = tmp
}

// RemoveQueue defined to remove an item from the heap
func RemoveQueue() {

}

// PollQueue defined to get the top element from the heap
func PollQueue() {

}

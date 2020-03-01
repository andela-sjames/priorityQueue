package main

import (
	"github.com/andela-sjames/priorityQueue/pqueue"
)

func main() {

	// Priority Queue ADT uses the HEAP implementation.
	// heapArr := []int{4, 7, 8, 3, 2, 6, 5, 1, 3, 5, 4, 6, 13, 10, 9, 8, 15, 17} // Non Heap Array

	pqArr := []pqueue.PItem{
		{"Visit China", 4},
		{"Visit Japan", 7},
		{"Visit China", 8},
		{"Visit Japan", 3},
		{"Visit China", 2},
		{"Visit Japan", 6},
		{"Visit China", 5},
		{"Visit Japan", 1},
		{"Visit China", 3},
		{"Visit Japan", 5},
		{"Visit China", 4},
		{"Visit Japan", 6},
		{"Visit China", 13},
		{"Visit Japan", 10},
		{"Visit China", 9},
		{"Visit Japan", 8},
		{"Visit China", 15},
		{"Visit Japan", 17},
	}

	pqueue.BuildHeap(pqArr, len(pqArr))
	pqueue.PrintHeap(pqArr, len(pqArr))
}

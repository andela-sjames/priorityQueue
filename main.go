package main

import (
	"github.com/andela-sjames/priorityQueue/pqueue"
)

func main() {

	// Priority Queue ADT uses the HEAP implementation.
	heapArr := []int{4, 7, 8, 3, 2, 6, 5, 1, 3, 5, 4, 6, 13, 10, 9, 8, 15, 17} // Non Heap Array

	pqueue.BuildHeap(heapArr, len(heapArr))
	pqueue.PrintHeap(heapArr, len(heapArr))
}

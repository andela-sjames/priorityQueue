package main

import (
	"github.com/andela-sjames/priorityQueue/pqueue"
)

func main() {

	// Priority Queue ADT uses the HEAP Invariant implementation.
	// heapArr := []int{4, 7, 8, 3, 2, 6, 5, 1, 3, 5, 4, 6, 13, 10, 9, 8, 15, 17} // Non Heap Array

	// pqArr := []pqueue.PItem{
	// 	{Item: "Visit China", Priority: 4},
	// 	{Item: "Visit Japan", Priority: 7},
	// 	{Item: "Eat Pizza", Priority: 8},
	// 	{Item: "Run marathon", Priority: 3},
	// 	{Item: "Buy Banana", Priority: 2},
	// 	{Item: "Go Fishing", Priority: 6},
	// 	{Item: "Skipping", Priority: 5},
	// 	{Item: "Swimming", Priority: 1},
	// 	{Item: "Go Kart", Priority: 3},
	// 	{Item: "Ice Skating", Priority: 5},
	// 	{Item: "Do Kung Fu", Priority: 4},
	// 	{Item: "Read Book", Priority: 6},
	// 	{Item: "Watch Netflix", Priority: 13},
	// 	{Item: "Get a boyfriend", Priority: 10},
	// 	{Item: "Visit France", Priority: 9},
	// 	{Item: "Study Espanol", Priority: 8},
	// 	{Item: "Buy a country", Priority: 15},
	// 	{Item: "Sahre the money", Priority: 17},
	// }

	pqArr2 := []pqueue.PItem{
		{Item: "Visit China", Priority: 4},
		{Item: "Visit Japan", Priority: 7},
	}

	pqueue.BuildHeap(pqArr2, len(pqArr2))
	pqueue.PrintHeap(pqArr2, len(pqArr2))
}

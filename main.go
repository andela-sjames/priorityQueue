package main

import (
	"fmt"

	"github.com/andela-sjames/priorityQueue/pqueue"
)

func main() {

	// Priority Queue ADT uses the HEAP Invariant implementation.
	maxheap := pqueue.NewHeap()
	maxheap.InsertPriority("Visit China", 4)
	maxheap.InsertPriority("Visit Japan", 7)
	maxheap.InsertPriority("Eat Pizza", 8)
	maxheap.InsertPriority("Run marathon", 3)
	maxheap.InsertPriority("Buy Banana", 2)
	maxheap.InsertPriority("Go Fishing", 6)
	maxheap.InsertPriority("Skipping", 5)
	maxheap.InsertPriority("Swimming", 1)
	maxheap.InsertPriority("Go Kart", 3)
	maxheap.InsertPriority("Ice Skating", 5)
	maxheap.InsertPriority("Do Kung Fu", 4)
	maxheap.InsertPriority("Watch Netflix", 13)
	maxheap.InsertPriority("Get a boyfriend", 10)
	maxheap.InsertPriority("Visit France", 21)
	maxheap.InsertPriority("Study Espanol", 8)
	maxheap.InsertPriority("Buy a country", 15)
	maxheap.InsertPriority("Share the money", 17)

	item, _ := maxheap.ShowPriority()
	fmt.Println(item, "The priority")

	// maxheap.PrintHeap()
}

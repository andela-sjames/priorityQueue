package main

import (
	"fmt"

	"github.com/andela-sjames/priorityQueue/pqueue"
)

func main() {

	// Priority Queue ADT uses the HEAP Invariant implementation.
	maxheap := pqueue.NewHeap()
	maxheap.InsertPriority("Visit China", 2)
	maxheap.InsertPriority("Visit Japan", 7)
	maxheap.InsertPriority("Eat Pizza", 2)
	maxheap.InsertPriority("Run marathon", 11)
	maxheap.InsertPriority("Buy Banana", 7)
	maxheap.InsertPriority("Go Fishing", 13)
	maxheap.InsertPriority("Skipping", 2)

	maxheap.InsertPriority("Swimming", 2)
	maxheap.InsertPriority("Go Kart", 3)
	maxheap.InsertPriority("Ice Skating", 5)
	maxheap.InsertPriority("Do Kung Fu", 4)
	maxheap.InsertPriority("Watch Netflix", 13)
	maxheap.InsertPriority("Get a boyfriend", 10)
	maxheap.InsertPriority("Visit France", 21)
	maxheap.InsertPriority("Study Espanol", 8)
	maxheap.InsertPriority("Buy a country", 15)
	maxheap.InsertPriority("Share the money", 17)

	item, p := maxheap.ShowPriority()
	fmt.Println(item, p, "The priority")

	// s := maxheap.ShowHeap()
	// fmt.Println(s)

	t := maxheap.ShowHashTable()
	fmt.Println(t)

	item, p = maxheap.Poll()
	fmt.Println(item, p, "The poll")

	t = maxheap.ShowHashTable()
	fmt.Println(t)

	_ = maxheap.Remove(11)

	maxheap.ShowHashTable()
}

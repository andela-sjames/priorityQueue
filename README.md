# priorityQueue

Implementation of the Priority Queue Abstract Data Type (ADT) using the Binary Heap Invariant.

## QUICK NOTE

In addition to the Binary Heap Implementation, this project makes use of a Hash table which gives us a time complexity
of O(1) time for search in comparison to 0(n) time for search without a Hash table.

An ADT simply defines the behaviour that a DataType(priorityQueue) should have; the Binary Heap Invariant gives
us a better time complexity than using a linked list or an unsorted list.

Usage: <https://play.golang.org/p/BTO5XqSFghX>

```go

package main

import (
    "fmt"
    pqueue "github.com/andela-sjames/priorityQueue"
)

func main() {
    maxheap := pqueue.NewHeap()

    maxheap.InsertPriority("Visit China", 2)
    maxheap.InsertPriority("Visit Japan", 7)
    maxheap.InsertPriority("Eat Pizza", 2)
    maxheap.InsertPriority("Run marathon", 11)
    maxheap.InsertPriority("Buy Banana", 7)
    maxheap.InsertPriority("Go Fishing", 13)
    maxheap.InsertPriority("Skipping", 2)

    item, p := maxheap.ShowPriority()
    fmt.Println(item, p, "The priority") // Go Fishing 13 The priority

    s := maxheap.ShowHeap()
    fmt.Println(s)
    // [{"Item":"Go Fishing","Priority":13} {"Item":"Visit Japan","Priority":7} {"Item":"Run marathon","Priority":11} {"Item":"Visit China","Priority":2} {"Item":"Buy Banana","Priority":7} {"Item":"Eat Pizza","Priority":2} {"Item":"Skipping","Priority":2}]

    t := maxheap.ShowHashTable()
    fmt.Println(t) // &map[2:[3 5 6] 7:[1 4] 11:[2] 13:[0]]

    item, p = maxheap.Poll()
    fmt.Println(item, p, "The poll") // Go Fishing 13 The poll

    t = maxheap.ShowHashTable()
    fmt.Println(t) // &map[2:[3 5 2] 7:[1 4] 11:[0]]

    _ = maxheap.Remove(11)

    t = maxheap.ShowHashTable()
    fmt.Println(t) // &map[2:[3 2 4] 7:[0 1]]

    l := maxheap.Length()
    fmt.Println(t) // 5

    // Set Min option is to true for minheap
    minheap := pqueue.NewHeap(pqueue.Options{
        Min: true,
    })

}
```

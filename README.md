# priorityQueue

Implementation of the Priority Queue ADT using Binary Heap Invariant

Usage:

```go

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
```

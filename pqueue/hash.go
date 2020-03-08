package pqueue

import (
	"fmt"
	"sync"
)

type htable map[int][]int

// HeapHash defined to create hash table for Heap data structure
// can be used for a min heap or max heap invariant.
type HeapHash struct {
	sync.RWMutex
	m htable
}

// NewHeapHash defined to create a new
// HeapHash object
func NewHeapHash() *HeapHash {
	h := &HeapHash{}
	h.m = make(htable)
	return h
}

// AddToTable struct defined
// If the value does not exist, add a new index
// to the array and make a new key, pair
// otherwise, get the exiting value(array),
// update it (to avoid collision) and update the key, pair
// i.e. key = updated_array_value
func (h *HeapHash) AddToTable(priority, index int) {

	arr := make([]int, 0)

	h.Lock()
	defer h.Unlock()

	if val, exist := h.m[priority]; !exist {
		arr = append(arr, index)
		h.m[priority] = arr
	} else {
		// avoid repeated entry
		if ok := h.intInSlice(index, val); !ok {
			val = append(val, index)
			h.m[priority] = val
		}

	}
}

// GetFromTable struct defined
// returns the index of the priority
func (h *HeapHash) GetFromTable(priority int, index int) int {
	var result int

	h.Lock()
	defer h.Unlock()

	if val, exist := h.m[priority]; exist {
		if len(val) == 1 {
			result = val[0]
			h.m[priority] = nil
		} else {
			// get the value x from the slice
			var idx int
			var arr []int
			idx, arr = h.removeFromSlice(index, val)
			if idx == index {
				h.m[priority] = arr
				result = idx
			}

		}
	}
	return result
}

// PeekPriority uses the priority key given
// to return the first item of the value
// but does not delete from the hash table
func (h *HeapHash) PeekPriority(priority int) int {

	h.Lock()
	defer h.Unlock()

	if val, exist := h.m[priority]; exist {
		return val[0]
	}
	return -1
}

// DeleteFromTable struct defined
func (h *HeapHash) DeleteFromTable(priority int) {

	h.Lock()
	defer h.Unlock()

	if _, exist := h.m[priority]; exist {
		delete(h.m, priority)
	}
}

// ShowHashTable show the content of the hash table
func (h *HeapHash) ShowHashTable() {
	fmt.Println(h.m)
}

func (h *HeapHash) intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (h *HeapHash) removeFromSlice(itm int, list []int) (int, []int) {
	for i := range list {
		if list[i] == itm {
			// get&remove itm from slice
			list = append(list[:i], list[i+1:]...)
			return itm, list
		}
	}
	return -1, list
}

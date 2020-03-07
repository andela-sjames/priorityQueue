package pqueue

import (
	"fmt"
	"sync"
)

type htable map[int][]int

// Pop from queue
// x, a = a[0], a[1:]

// Pop from stack
// x, a = a[len(a)-1], a[:len(a)-1]

// Push
// a = append(a, x)

// HashTable struct defined
var HashTable = struct {
	sync.RWMutex
	m htable
}{m: make(htable)}

// AddToTable struct defined
// If the value does not exist, add a new index
// to the array and make a new key, pair
// otherwise, get the exiting value(array),
// update it (to avoid collision) and update the key, pair
// i.e. key = updated_array_value
func AddToTable(priority, index int) {

	arr := make([]int, 0)

	HashTable.Lock()
	defer HashTable.Unlock()

	if val, exist := HashTable.m[priority]; !exist {
		arr = append(arr, index)
		HashTable.m[priority] = arr
	} else {
		// avoid repeated entry
		if ok := intInSlice(index, val); !ok {
			val = append(val, index)
			HashTable.m[priority] = val
		}

	}
}

// GetFromTable struct defined
// returns the index of the priority
func GetFromTable(priority int, index int) int {
	var result int

	HashTable.Lock()
	defer HashTable.Unlock()

	if val, exist := HashTable.m[priority]; exist {
		if len(val) == 1 {
			result = val[0]
			HashTable.m[priority] = nil
		} else {
			// get the value x from the slice
			var idx int
			var arr []int
			idx, arr = removeFromSlice(index, val)
			if idx == index {
				HashTable.m[priority] = arr
				result = idx
			}

		}
	}
	return result
}

// PeekPriority uses the priority key given
// to return the first item of the value
// but does not delete from the hash table
func PeekPriority(priority int) int {

	HashTable.Lock()
	defer HashTable.Unlock()

	if val, exist := HashTable.m[priority]; exist {
		return val[0]
	}
	return -1
}

// DeleteFromTable struct defined
func DeleteFromTable(priority int) {

	HashTable.Lock()
	defer HashTable.Unlock()

	if _, exist := HashTable.m[priority]; exist {
		delete(HashTable.m, priority)
	}
}

// ShowHashTable show the content of the hash table
func ShowHashTable() {
	fmt.Println(HashTable.m)
}

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func removeFromSlice(itm int, list []int) (int, []int) {
	for i := range list {
		if list[i] == itm {
			// get&remove itm from slice
			list = append(list[:i], list[i+1:]...)
			return itm, list
		}
	}
	return -1, list
}

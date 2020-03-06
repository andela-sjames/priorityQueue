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
		if ok := intInSlice(index, val); !ok {
			val = append(val, index)
		}
		HashTable.m[priority] = val
	}
}

// GetFromTable struct defined
// returns the index of the priority
func GetFromTable(priority int) int {
	var result int
	var x int

	HashTable.Lock()
	defer HashTable.Unlock()

	if val, exist := HashTable.m[priority]; exist {
		if len(val) == 1 {
			result = val[0]
			val = nil
		} else {
			x, val = val[0], val[1:]
			HashTable.m[priority] = val
			result = x
		}
	}
	return result
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

package pqueue

import (
	"sync"
)

type htable map[int][]int

// HashTable struct defined
var hashTable = struct {
	sync.RWMutex
	m htable
}{m: make(htable)}

// AddToTable struct defined
func AddToTable(priority, index int) {

	/**
	If the value does not exist, add an new index
	to the array and make a new key, pair
	otherwise, get the exiting value(array),
	update it (to avoid collision) and update the key, pair
	i.e. key = updated_array_value
	*/
	arr := make([]int, 0)

	hashTable.Lock()
	defer hashTable.Unlock()

	if val, exist := hashTable.m[priority]; !exist {
		arr = append(arr, index)
		hashTable.m[priority] = arr
	} else {
		val = append(val, index)
		hashTable.m[priority] = val
	}
}

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

// InsertTable struct defined
func InsertTable(priority, index int) {

	// val == array, key == priority
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

package pqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPriorityQueue(t *testing.T) {
	maxheap := NewHeap()
	assert := assert.New(t)

	maxheap.InsertPriority("Visit China", 2)
	maxheap.InsertPriority("Visit Japan", 7)
	maxheap.InsertPriority("Eat Pizza", 2)
	maxheap.InsertPriority("Run marathon", 11)

	t.Run("Show priority returns item and priority", func(t *testing.T) {
		item, p := maxheap.ShowPriority()
		assert.Equal(item, "Run marathon")
		assert.Equal(p, 11)

	})

	t.Run("Poll an object from the queue", func(t *testing.T) {

	})

	t.Run("Remove an object from the heap", func(t *testing.T) {

	})
}

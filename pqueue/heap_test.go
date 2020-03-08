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

	t.Run("ShowPriority should return item and priority", func(t *testing.T) {
		item, p := maxheap.ShowPriority()
		assert.Equal(item, "Run marathon")
		assert.Equal(p, 11)

	})

	t.Run("Poll an object from the queue", func(t *testing.T) {
		maxheap.InsertPriority("Watch Netflix", 13)
		maxheap.InsertPriority("Get a boyfriend", 10)
		maxheap.InsertPriority("Visit France", 21)

		item, p := maxheap.Poll()
		assert.Equal(item, "Visit France")
		assert.Equal(p, 21)

	})

	t.Run("Remove an object from the heap", func(t *testing.T) {

		l := maxheap.Length()
		assert.Equal(l, 6)

		done := maxheap.Remove(11)

		l = maxheap.Length()
		assert.Equal(done, true)
		assert.Equal(l, 5)
	})
}

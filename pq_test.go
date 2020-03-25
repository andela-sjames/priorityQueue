package pqueue_test

import (
	"testing"

	pqueue "github.com/andela-sjames/priorityQueue"

	"github.com/stretchr/testify/assert"
)

func TestMaxPriorityQueue(t *testing.T) {
	// defaults to max heap if Min option is
	// not set to true
	maxheap := pqueue.NewHeap(pqueue.Options{})
	assert := assert.New(t)

	maxheap.InsertPriority("Visit China", 2)
	maxheap.InsertPriority("Visit Japan", 7)
	maxheap.InsertPriority("Eat Pizza", 2)
	maxheap.InsertPriority("Run marathon", 11)
	maxheap.InsertPriority("Go Kart", 3)
	maxheap.InsertPriority("Ice Skating", 5)
	maxheap.InsertPriority("Do Kung Fu", 4)
	maxheap.InsertPriority("Get a boyfriend", 10)

	t.Run("ShowPriority should return item and priority", func(t *testing.T) {
		item, p := maxheap.ShowPriority()
		assert.Equal("Run marathon", item)
		assert.Equal(11, p)

	})

	t.Run("Poll an object from the queue", func(t *testing.T) {
		maxheap.InsertPriority("Watch Netflix", 13)
		maxheap.InsertPriority("Get a boyfriend", 10)
		maxheap.InsertPriority("Visit France", 21)

		item, p := maxheap.Poll()
		assert.Equal("Visit France", item)
		assert.Equal(21, p)

	})

	t.Run("Remove an object from the heap", func(t *testing.T) {

		l := maxheap.Length()
		assert.Equal(10, l)

		done := maxheap.Remove(11)

		l = maxheap.Length()
		assert.Equal(true, done)
		assert.Equal(9, l)
	})

	t.Run("Heap table can be returned", func(t *testing.T) {
		v := maxheap.ShowHashTable()
		assert.NotNil(v)

	})

	t.Run("Heap content can be shown", func(t *testing.T) {
		s := maxheap.ShowHeap()
		assert.NotNil(s)
	})
}

func TestMinPriorityQueue(t *testing.T) {

	// Set Min option is to true for minheap
	minheap := pqueue.NewHeap(pqueue.Options{
		Min: true,
	})
	assert := assert.New(t)

	minheap.InsertPriority("Visit China", 2)
	minheap.InsertPriority("Visit Japan", 7)
	minheap.InsertPriority("Eat Pizza", 2)
	minheap.InsertPriority("Run marathon", 11)
	minheap.InsertPriority("Go Kart", 3)
	minheap.InsertPriority("Ice Skating", 5)
	minheap.InsertPriority("Do Kung Fu", 4)
	minheap.InsertPriority("Get a boyfriend", 10)

	t.Run("ShowPriority should return item and priority", func(t *testing.T) {
		item, p := minheap.ShowPriority()
		assert.Equal("Visit China", item)
		assert.Equal(2, p)

	})

	t.Run("Poll an object from the queue", func(t *testing.T) {
		minheap.InsertPriority("Watch Netflix", 13)
		minheap.InsertPriority("Get a boyfriend", 10)
		minheap.InsertPriority("Visit France", 21)

		item, p := minheap.Poll()
		assert.Equal("Visit China", item)
		assert.Equal(2, p)

	})

	t.Run("Remove an object from the heap", func(t *testing.T) {

		l := minheap.Length()
		assert.Equal(10, l)

		done := minheap.Remove(11)

		l = minheap.Length()
		assert.Equal(true, done)
		assert.Equal(9, l)
	})

	t.Run("Heap table can be returned", func(t *testing.T) {
		v := minheap.ShowHashTable()
		assert.NotNil(v)

	})

	t.Run("Heap content can be shown", func(t *testing.T) {
		s := minheap.ShowHeap()
		assert.NotNil(s)
	})
}

package main

import "fmt"

func main() {
	fmt.Println("Define the tree")

	arr := []int{2, 3, 4, 7, 5, 8, 9, 8, 10, 12, 11, 13, 8}

	constructHeap(arr)
}

type node struct {
	key   int
	left  interface{}
	right interface{}
}

func constructHeap(arr []int) {

	for i := range arr {
		heapify(i, arr)
	}
}

func heapify(i int, arr []int) {

	var rootNode node

	left := 2*i + 1
	right := 2*i + 2

	rootNode = node{key: arr[i]}
	rootNode.left = node{key: arr[left]}
	rootNode.right = node{key: arr[right]}

	heapify(i, arr)
}

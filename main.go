package main

import "fmt"

func main() {
	fmt.Println("Define the tree")

	arr := []int{2, 3, 4, 7, 5, 8, 9, 8, 10, 12, 11, 13, 8}
	var root interface{}

	l := len(arr)
	// root := interface{}(nil)

	root = constructBTree(arr, root, 0, l)
	fmt.Println(root)

}

type node struct {
	key   int
	left  interface{}
	right interface{}
}

// let's construct a simple Binary tree
func constructBTree(arr []int, root interface{}, index int, l int) interface{} {

	var leftIndex int
	var rigntIndex int

	// base condition
	if index < l {
		n := node{key: arr[index]}
		root := n

		leftIndex = 2*index + 1
		rigntIndex = 2*index + 2

		constructBTree(arr, root.left, leftIndex, l)
		constructBTree(arr, root.right, rigntIndex, l)
	}

	return root
}

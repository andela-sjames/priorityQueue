package queue

import "fmt"

func main() {
	fmt.Println("Define the tree")
}

func generateBaseNode() {

}

type node struct {
	name  string
	left  interface{}
	right interface{}
}

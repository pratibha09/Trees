package main

import (
	"fmt"
)

type node struct {
	data int
	left *node
	right *node
}

func insert(head *node, data int) {
	newNode := &node{data: data, left : nil, right : nil}
	temp := head

	for {
		if data < temp.data {
			if temp.left == nil {
				temp.left = newNode
				return
			}
			temp = temp.left

		}else if data >= temp.data {

			if temp.right == nil {
				temp.right = newNode
				return
			}
			temp = temp.right
		}
	}
}

func size(n *node) int{
	if n == nil {
		return 0
	}else {
		return (size(n.left)+size(n.right)+1)
		
	}
}

func main() {
	var head *node
	head = &node{data: 3, left: nil, right: nil}
	insert(head, 2)
	insert(head, 1)
	insert(head, 4)
	insert(head, 5)
	insert(head, 6)	
	fmt.Println("size of tree", size(head))

}



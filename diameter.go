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

func diameter(n *node) int {

	if n == nil {
		return 0
	}

	lheight := height(n.left)
	rheight := height(n.right) 
	
	ldiameter := diameter(n.left)
	rdiameter := diameter(n.right)

	return max(lheight + rheight + 1, max(ldiameter, rdiameter))
}

func height(node *node)int {
	if node == nil {
		return 0
	}else {
		return 1 + max(height(node.left), height(node.right))
	}
}
	
func max(a, b int) int{
	if a >= b {
		return a
	}else{
		return b
	}
}

func main(){
	var head *node
	head = &node{
		data:2,
		left: nil,
		right: nil,
		}
	insert(head, 1)
	insert(head, 3)
	insert(head, 4)
	insert(head, 5)
	insert(head, 6)
	insert(head, 7)
	insert(head, 8)
	diameter(head)
	fmt.Println("diameter", diameter(head))
}

//time complexity is o(n^2)


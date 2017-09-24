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

func depth(n *node) int{
	if n == nil {
		return 0
	}else {
		ldepth := depth(n.left)    //calc the depth of left subtree
		rdepth := depth(n.right)   //calc the depth of right subtree
		if ldepth > rdepth {
			return ldepth + 1
		}else {
			return rdepth + 1
		}
	}
}

func main(){
	var head *node
	head = &node{
		data:2,
		left: nil,
		right: nil,
		}
	insert(head, 3)
	insert(head, 1)
	insert(head, 5)
	insert(head, 6)
	insert(head, 7)
	fmt.Println("height of tree",depth(head))

}

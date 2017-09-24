package main

import(
	"fmt"
)

type node struct {
	data int
	left *node
	right *node
}

func getleafCount(n *node) int{
	if n == nil {
		return 0
	}else if n.left == nil && n.right == nil {
		return 1
	}else {
		return getleafCount(n.left) + getleafCount(n.right)
	}
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

func main(){
	var head *node
	head = &node{
		data : 4,
		left : nil,
		right : nil,
	}

	insert(head, 2)
	insert(head, 5)
	insert(head, 1)
	insert(head, 3)	

	getleafCount(head)
	fmt.Printf("leaf count of tree %d\n",getleafCount(head))
}



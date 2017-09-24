package main

import(
	"fmt"
)

type node struct {
	data int
	left *node
	right *node
}

func insert(head *node, data int) {
	newnode := &node{data : data, left : nil, right : nil}
	temp := head
	for {	
		if data < temp.data {
			if temp.left == nil {
				temp.left = newnode
				return
			}
			temp = temp.left
	
		}else if data >= temp.data {
			if temp.right == nil{
				temp.right = newnode
				return
			}
			temp =  temp.right
		}
	}
}

func minvalue(n *node) int {
	current := n
	for current.left != nil {
		current = current.left
	}
	return current.data
}

func main(){
	var head *node
	head = &node{
		left: nil,
		right: nil,
		data: 3,
		}
	insert(head, 1)
	insert(head, 2)
	insert(head, 1)
	insert(head, 8)
	insert(head,13)
	minvalue(head)
	fmt.Println("min value is", minvalue(head))
}




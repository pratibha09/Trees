package main

import (
	"fmt"
)

type node struct {
	data int
	left  *node
	right *node
}

func insert(head *node, data int) {
	newNode := &node{data : data, left: nil, right: nil}
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

func printrange(head *node, r1 int, r2 int) {
	if head == nil {
		return
	}
	if r1 < head.data {
		printrange(head.left, r1, r2)
	}
	if r1 <= head.data && r2 >= head.data {
		fmt.Printf("%d\n", head.data)
	} 
	if r2 > head.data {
		printrange(head.right, r1, r2)
	}
}

func main(){
	r1 := 10 
	r2 := 25
	var head *node
	head = &node{
		data: 20,
		left: nil,
		right: nil,
		}
	insert(head, 8)
	insert(head, 22)
	//insert(head, 1)
	insert(head, 4)
	insert(head, 12)
	//insert(head, 15)
	
	printrange(head, r1, r2)
}

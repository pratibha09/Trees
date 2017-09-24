package main

import (
	"fmt"
)

type node struct{
	data int
	left *node
	right *node
}

func sum(n *node)bool {
	var left_data = 0
	var right_data = 0

	if n == nil || n.left == nil && n.right == nil {
		return false
	}else{
		if n.left != nil {
			left_data = n.left.data
		} 
		if n.right != nil {
			right_data = n.right.data
		}
		if (n.data == left_data + right_data) && (sum(n.left) && sum(n.right)){
			return false
		}else {
			return true
		}
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
	if sum(head){
		fmt.Println("given children satify the sum property", sum(head))
	}else{
		fmt.Println("they dont!")
	}
}


//time complexity  is O(n)

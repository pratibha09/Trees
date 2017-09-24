package main

import(
	"fmt"
)
type node struct {
	data int
	left *node
	right *node
}

func insert(head *node, data int){
	newnode := &node{data: data, left : nil, right : nil}
	temp := head 
	for  {
		if data < temp.data {
			if temp.left == nil {
				temp.left = newnode
				return
			}
			temp = temp.left
		}else if data >= temp.data {
			if temp.right == nil {
				temp.right = newnode
				return
			}
			temp = temp.right
		}
	}

}


func balanced(head *node)bool {
	if head == nil {
		return false
	}
	lh := height(head.left)
	rh := height(head.right)
	if (lh-rh) <= 1 && balanced(head.left) && balanced(head.right) {
		return true
	}else{
		return false
	}
}


func  height(n *node) int {
	if n == nil {
		return 0
	}else {
		lh := height(n.left)
		rh := height(n.right)
		if lh > rh {
			return lh + 1
		}else {
			return rh + 1
		}
	}
}

func main() {
	head := &node{
		data : 1, 
		left : nil,
		right : nil,
		} 

	insert(head, 2)
	insert(head, 3)
	insert(head, 1)
	insert(head, 5)
	//insert(head, 7)
	//insert(head, 1)
	//insert(head, 8)
	balanced(head)
	if balanced(head){
		fmt.Println("is balanced")
	}else{
		fmt.Println("its not")
	}

}


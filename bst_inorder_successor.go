package main

import(
	"fmt"
)

type node struct{
	data int
	left *node
	right *node
	parent *node
}

func insert(head *node, data int){
	newnode := &node{data : data, left : nil, right : nil}
	temp := head
	for {
		if data < temp.data{
			if temp.left == nil{
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

func inordersuccessor(head *node, n *node) *node {
	if n.right != nil{
		return minvalue(n.right)
	}
	var p *node
	p = n.parent
	for p != nil && n == p.right {
		n = p
		p = p.parent
	}	
	return p
}

func minvalue(n *node) *node {
	var current *node
	current = n
	for current.left != nil {
		current = current.left
	}
	return current
}


func main(){
	var head *node
	head = &node{
		data : 20,
		left : nil,
		right : nil,	
		}

	insert(head, 8)
	insert(head, 22)
	insert(head, 4)
	insert(head, 12)
	insert(head, 10)
	insert(head, 14)
	temp := head.left.right
	succ := inordersuccessor(head, temp)
	if succ != nil {
		fmt.Printf("inorder successor of %d is %d\n", temp.data, succ.data)
	}else {
		fmt.Println("inorder successor doesnt exist")
	}
}

//time complexity is O(h) where h is height of tree




package main

import (
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

func lca (head *node, n1 int, n2 int) *node {
	if head == nil {
		return head
	}
	if head.data > n1 && head.data > n2 {
		return lca (head.left, n1, n2)
	}
	if head.data < n1 && head.data < n2 {
		return lca (head.right,n1 , n2)
	}
	return head
}

func main() {
	head := &node{
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

	n1 := 10; n2 := 14
	var p *node =  lca(head, n1, n2)
	fmt.Printf("LCA of  %d and %d is: %d\n", n1, n2, p.data)


}

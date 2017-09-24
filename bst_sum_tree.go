package main

import(
	"fmt"
)

type node struct {
	data int
	left  *node
	right *node
}

func insert(head *node, data int){
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
			temp = temp.right
		}
	}
}

func bstsumtree(head *node, sum *int) {
	if head == nil {
		return
	}
	bstsumtree(head.right, sum)
	*sum = *sum + head.data
	head.data = *sum - head.data
	bstsumtree(head.left, sum)
}

func transformtree(head *node){
	sum := 0
	bstsumtree(head, &sum)
}

func printorder(head *node){
	if head == nil {
		return
	}
	printorder(head.left)
	fmt.Println(head.data)
	printorder(head.right)
}

func main(){
	head := &node{
		data : 11,
		left : nil,
		right : nil,
		}
	insert(head, 2)
	insert(head, 29)
	insert(head, 1)
	insert(head, 7)
	insert(head, 15)
	insert(head, 35)
	fmt.Println("inorder traversal")
	printorder(head)
	transformtree(head)
	fmt.Println("transformed inorder")
	printorder(head)
}

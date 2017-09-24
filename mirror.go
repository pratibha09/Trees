package main

import(
	"fmt"
)

type node struct{
	data int
	left *node
	right *node
	temp *node
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

func mirror(n *node){
	if n == nil {
		return
	}else {
		temp := n.left
		n.left = n.right
		n.right = temp
		
		mirror(n.left)
		mirror(n.right)
	}
}

func inorder(n *node){
	if n == nil {
		return
	}else{
		inorder(n.left)
		fmt.Println("%d", n.data)
		inorder(n.right)
	}
}


func main(){
	var head *node
	head = &node{
		data:4,
		left: nil,
		right: nil,
		}
	insert(head, 3)
	insert(head, 2)
	insert(head, 5)
	insert(head, 6)
	insert(head, 7)

	println("inorder of actual tree")
	inorder(head)
	mirror(head)

	println("inorder of mirror tree")
	inorder(head)
	mirror(head)
	fmt.Println("mirror of tree")
}

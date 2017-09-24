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
	newnode := &node{left: nil, right: nil, data: data}
	temp := head

	for {
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

func printgivenlevel(head *node, level int){
	if head == nil {
		return
	}else if level == 1{
		fmt.Printf("%d\n", head.data)
	}else if level >1 {
		printgivenlevel(head.left, level-1)
		printgivenlevel(head.right, level-1)
	}
}

func printlevelorder (head *node){
	h := height(head)
	for i := 0; i<= h; i++ {
		printgivenlevel(head, i)
	}
}

func height (n *node) int {
	if n == nil {
		return 0
	} else {
		lheight := height (n.left)
		rheight := height (n.right)
		if lheight > rheight {
			return lheight + 1
		} else {
			return rheight + 1
		}
	}
}

func main() {
	var head *node
	head = &node{
		left: nil,
		right: nil,
		data: 1,
		}
	insert(head, 2)
	insert(head, 3)
	insert(head, 4)
	insert(head, 5)
	insert(head,16)

	fmt.Printf("order traversal of binary tree \n")
	height(head)
	printlevelorder(head)
	
}




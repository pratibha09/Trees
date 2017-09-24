package main

import(
	"fmt"
)

type node struct {
	data int
	left *node
	right *node
}

func Delete(n *node){
	if n == nil {
		return
	}
	Delete(n.left)
	Delete(n.right)
	fmt.Println("deleting node with data", n.data)

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
		data:2,
		left: nil,
		right: nil,
		}
	insert(head, 3)
	insert(head, 4)
	insert(head, 5)
	insert(head, 6)
	insert(head, 7)

	Delete(head)
	head = nil
	fmt.Println("tree deleted")

}


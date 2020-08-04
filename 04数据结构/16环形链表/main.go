package main

import "fmt"

type CatNode struct {
	no int
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head
		fmt.Println(newCatNode,"加入到环形链表")
		return
	}
	temp := head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newCatNode
	newCatNode.next = head
}

func ListCatNode(head *CatNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("kong")
		return
	}
	for temp.next != nil {
		fmt.Printf("[%d %s]",temp.no,temp.name)
		if temp.next == head {break}
		temp = temp.next
	}

}

func DelCatNode(head *CatNode,id int)
func main() {

}

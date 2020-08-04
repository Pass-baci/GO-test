package main

import (
	"fmt"
)

type HeroNode struct {
	no int
	name string
	pre *HeroNode
	next *HeroNode
}

func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newHeroNode
	newHeroNode.pre = temp
}

func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	flag := true
	for temp.next != nil {
		if temp.next.no > newHeroNode.no {
			break
		} else if temp.next.no == newHeroNode.no {
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("没有")
		return
	} else {
		newHeroNode.next = temp.next
		newHeroNode.pre = temp
		if temp.next != nil {
			temp.next.pre = newHeroNode
		}
		temp.next = newHeroNode
	}

}

func DelHeroNode(head *HeroNode, id int) {
	temp := head
	flag := false
	for temp.next != nil {
		if temp.next.no == id {
			flag = true
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("meiyou")
		return
	} else {
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		}
	}
}

func ListHeroNode(head *HeroNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("kong")
		return
	}
	for temp.next != nil {
		fmt.Printf("[%d,%s]==>",temp.next.no,temp.next.name)
		temp = temp.next
	}
}

func ListHeroNode2(head *HeroNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("kong")
		return
	}
	for temp.next != nil {
		temp = temp.next
	}
	for temp.pre != nil {
		fmt.Printf("[%d,%s]==>",temp.no,temp.name)
		temp = temp.pre
	}
}

func main(){
	head := &HeroNode{}
	hero1 := &HeroNode{
		no: 1,
		name: "song",
	}
	hero2 := &HeroNode{
		no: 2,
		name: "lu",
	}
	hero3 := &HeroNode{
		no: 3,
		name: "lin",
	}
	InsertHeroNode(head,hero1)
	InsertHeroNode(head,hero3)
	InsertHeroNode2(head,hero2)
	//DelHeroNode(head,2)
	//ListHeroNode(head)
	ListHeroNode2(head)
}

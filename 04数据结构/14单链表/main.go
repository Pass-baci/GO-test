package main

import "fmt"

type HeroNode struct {
	no int
	name string
	next *HeroNode
}

func InsertHeroNode(head *HeroNode,newHeroNode *HeroNode){
	temp := head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newHeroNode
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
		fmt.Println("已经存在no=",newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	}
}

func LisetHeroNOde(head *HeroNode){
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
	if flag {
		temp.next = temp.next.next
	} else {
		fmt.Println("bucunzai")
	}
}

func main() {
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
	DelHeroNode(head,2)
	LisetHeroNOde(head)
}

package main

import (
	"fmt"
	"os"
)

type Emp struct {
	id int
	name string
	next *Emp
}

type Emplink struct {
	head *Emp
}

type hashtable struct {
	arrlink [7]Emplink
}

func (e *Emplink)Insertemp(id int,name string) {
	var pre *Emp = nil
	cur := e.head
	temp := &Emp{
		id: id,
		name: name,
	}
	if cur == nil {
		e.head = temp
		return
	}
	if cur.id > id {
		temp1 := e.head
		e.head = temp
		temp.next = e.head.next
		e.head.next = temp1
		return
	}
	for cur != nil || cur.id < id {
		pre = cur
		cur = cur.next
	}
	pre.next = temp
	temp.next = cur
}

func (e *Emplink)ShowAllEmp(no int) {
	temp := e.head
	if temp == nil {
		fmt.Printf("第%d链表为空\n",no)
		return
	}
	for temp != nil {
		fmt.Printf("第%d链表,id=%d,name=%s => ",no,temp.id,temp.name)
		temp = temp.next
	}
	fmt.Println()
}

func (e *Emplink)SelectEmp(id int) {
	temp := e.head
	for {
		if temp == nil {
			fmt.Printf("没有查找到id = %d的数据",id)
			break
		}
		if temp.id == id {
			fmt.Printf("id = %d;name = %s",temp.id,temp.name)
			break
		}
		temp = temp.next
	}
}

func (e *Emplink)DeleteEmp(id int) {
	var pre *Emp = nil
	temp := e.head
	if temp.id == id {
		e.head = e.head.next
		return
	}
	for {
		if temp == nil {
			fmt.Printf("没有查找到id = %d的数据,无法删除",id)
			break
		}
		if temp.id == id {
			pre.next = temp.next
			fmt.Printf("已删除id = %d的数据",id)
			break
		}
	}
}

func (e *Emplink)ModifyEmp(id int,name string) {
	temp := e.head
	for {
		if temp == nil {
			fmt.Printf("没有查找到id = %d的数据,无法修改",id)
			break
		}
		if temp.id == id {
			temp.name = name
			fmt.Printf("已修改id = %d的数据,name = %s",temp.id,temp.name)
			break
		}
	}
}

func (h *hashtable)Inserthashtable(id int,name string){
	hashId := hashVal(id)
	h.arrlink[hashId].Insertemp(id,name)
}

func (h *hashtable)ShowHashTable() {
	for i := 0; i < len(h.arrlink); i++ {
		h.arrlink[i].ShowAllEmp(i+1)
	}
}

func (h *hashtable)SelectHashTable(id int) {
	val := hashVal(id)
	h.arrlink[val].SelectEmp(id)
}

func (h *hashtable)DeleteHashTable(id int) {
	val := hashVal(id)
	h.arrlink[val].DeleteEmp(id)
}
func (h *hashtable)ModifyHashTable(id int, name string) {
	val1 := hashVal(id)
	h.arrlink[val1].ModifyEmp(id, name)
}

func hashVal(id int) int {
	return id % 7
}

func main() {
	hashtabledemo := hashtable{}
	for {
		fmt.Println("请输入需要操作的序号：")
		fmt.Println("1.find")
		fmt.Println("2.ADD")
		fmt.Println("3.Modify")
		fmt.Println("4.Delete")
		fmt.Println("5.Show")
		fmt.Println("6.Exit")
		key := 0
		ID := 0
		name := ""
		fmt.Scan(&key)
		switch key {
		case 1:
			fmt.Println("请输入需要查找的id号:")
			fmt.Scan(&ID)
			hashtabledemo.SelectHashTable(ID)
		case 2:
			fmt.Println("请输入需要增加的id号:")
			fmt.Scan(&ID)
			fmt.Println("请输入需要增加的name:")
			fmt.Scan(&name)
			hashtabledemo.Inserthashtable(ID,name)
		case 3:
			fmt.Println("请输入需要修改的id号:")
			fmt.Scan(&ID)
			hashtabledemo.SelectHashTable(ID)
			fmt.Println("请输入需要修改的name:")
			fmt.Scan(&name)
			hashtabledemo.ModifyHashTable(ID,name)
		case 4:
			fmt.Println("请输入需要删除的id号:")
			fmt.Scan(&ID)
			hashtabledemo.DeleteHashTable(ID)
		case 5:
			hashtabledemo.ShowHashTable()
		case 6:
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入！")
		}
	}
}

package main

import "fmt"

type Boy struct {
	no int
	next *Boy
}

func AddBoy(num int) *Boy {
	first := &Boy{}
	curBoy := &Boy{}
	if num < 1 {
		fmt.Println("failed")
		return first
	}
	for i := 1; i <= num; i++ {
		boy := &Boy{
			no: i,
		}
		if i == 1 {
			first = boy
			curBoy = boy
			curBoy.next = first
		} else {
			curBoy.next = boy
			curBoy = boy
			curBoy.next = first
		}
	}
	return  first
}

func ShowBoy(first *Boy) {
	if first.next == nil {
		fmt.Println("kong")
		return
	}
	head := first
	for  {
		fmt.Printf("编号：[%d]\n",head.no)
		if head.next == first {break}
		head = head.next

	}
}

func PlayGame(first *Boy, startNo int,countNo int) {
	if first.next == nil {
		fmt.Println("kong")
		return
	}
	tail := first
	j := 1
	for tail.next != first {
		tail = tail.next
		j++
	}
	if j < startNo {
		fmt.Println("meiyoucunzai")
		return
	}
	for i:=1; i<=startNo-1; i++ {
		first = first.next
		tail = tail.next
	}
	for tail != first {
		for i := 1; i <= countNo-1; i++ {
			first = first.next
			tail = tail.next
		}
		fmt.Printf("编号【%d】出圈\n",first.no)
		first = first.next
		tail.next = first
	}
	fmt.Printf("编号【%d】出圈\n",first.no)

}
func main() {
	boy := AddBoy(100)
	ShowBoy(boy)
	PlayGame(boy,3,10)
}

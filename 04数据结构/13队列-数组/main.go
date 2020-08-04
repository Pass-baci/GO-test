package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

type Queue struct {
	maxSize int
	array [5]int
	front int
	rear int
}

func (q *Queue)AddQueue(val int)(err error) {
	if q.rear == q.maxSize-1 {
		err = errors.New("此队列已满")
		return err
	}
	q.rear++
	q.array[q.rear] = val
	fmt.Println("加入队列成功")
	return
}

func (q *Queue)GetQueue()(val int,err error) {
	if q.front == q.rear {
		err = errors.New("此队列为空")
		return
	}
	q.front++
	val = q.array[q.front]
	return val, nil
}

func (q *Queue)ShowQueue()(err error){
	if q.front == q.rear {
		err = errors.New("此队列为空")
		return err
	}
	for i := q.front+1; i <= q.rear; i++ {
		fmt.Printf("num:%d;val:%d\n",i,q.array[i])
	}
	return
}

func main() {
	queueNew := Queue{
		maxSize: 5,
		front: -1,
		rear: -1,
	}
	for {
		fmt.Println("请输入你所需要操作的序号")
		fmt.Println("1.Add;2.Get;3.Show;4.Exit")
		var KEY int
		fmt.Scan(&KEY)
		switch KEY {
		case 1:
			var val int
			fmt.Print("请输入需要加入队列的值：")
			fmt.Scan(&val)
			err := queueNew.AddQueue(val)
			if err != nil {
				fmt.Println(err)
			}
		case 2:
			queue, err := queueNew.GetQueue()
			fmt.Printf("出队列的值为：%d\n", queue)
			if err != nil {
				fmt.Println(err)
			}
		case 3:
			err := queueNew.ShowQueue()
			if err != nil {
				fmt.Println(err)
			}
		case 4:
			os.Exit(1)
		default:
			fmt.Println("输入有误")
		}
	}
}

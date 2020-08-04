package main

import "fmt"

func InsertSort(arr *[7]int){
	for i := 0; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex + 1] = arr[insertIndex]
			insertIndex --
		}
		if insertIndex + 1 != i {
			arr[insertIndex + 1] = insertVal
		}
		fmt.Printf("第%d次结果：%v\n",i+1,arr)
	}
}

func main() {
	arr := &[7]int{1,4,0,5,6,3,4}
	InsertSort(arr)
}

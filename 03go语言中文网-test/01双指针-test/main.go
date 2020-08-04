package main

import "fmt"

func sortArr(arr []int) []int{
	if len(arr) == 0 {
		return arr
	}
	b := 0
	for a := 0; a < len(arr); a++ {
		if arr[a] != 0 {
			arr[b] = arr[a]
			b++
		}
	}
	for i := b; i < len(arr); i++ {
		arr[b] = 0
	}
	return arr
}

func main() {
	a := []int{1,0,5,2,4,2,2}
	fmt.Println(sortArr(a))
}

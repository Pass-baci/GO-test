package main

import "github.com/golang/go/src/pkg/fmt"

func twoSum(nums []int, target int) []int {
	result := []int{}
	m := make(map[int]int)
	for i ,v := range nums {
		if v1, ok := m[target-v]; ok {
			result = append(result,v1,i)
		}
		m[v] = i
	}
	return result
}

func main() {
	nums := []int{2,2,7,7}
	sum := twoSum(nums, 9)
	fmt.Println(sum)
}

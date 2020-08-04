package main

import (
	"github.com/golang/go/src/pkg/fmt"
	"github.com/golang/go/src/pkg/sort"
)

func Solution(nums []int) [][]int {
	sort.Ints(nums)
	fmt.Println(nums)
	j := 0
	sum := nums[0]
	nums0 := [][]int{}
	for i := 1; i < j+3; i++ {
		sum +=nums[i]
		if sum == 0 {
			tempnums := nums[j:i+1]
			nums0 = append(nums0,tempnums)
		}
		if sum > 0 {break}
		if i == j+2 {
			j++
			if j == len(nums)-2 {break}
			i = j + 1
			sum = nums[j]
		}

	}
	return nums0
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	solution := Solution(nums)
	fmt.Println(solution)
}

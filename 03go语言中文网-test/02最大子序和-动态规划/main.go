package main

import "fmt"

func  maxSubArray(nums []int) int {
	ans := nums[0]
	sum := 0
	for _, v := range nums {
		if sum > 0 {
			sum += v
		} else {
			sum = v
		}
		ans = max(ans,sum)
	}
	return ans
}

func max(ans,sum int) int {
	if ans > sum {
		return ans
	} else {
		return sum
	}
}

func main() {
	a := []int{-2,1}
	num := maxSubArray(a)
	fmt.Println(num)
}

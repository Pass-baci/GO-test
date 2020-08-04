package main

import "fmt"

func findUnsortedSubarray(nums []int) int {
	l := len(nums)
	if l < 2  {
		return 0
	}
	low, high, min, max := 0, l-1, nums[0], nums[l-1]
	for i := 1; i < l; i++ {
		if min < nums[i] && i == low + 1 {
			low = i
			min = nums[i]
		}
		if max > nums[l-1-i] && high == l - i {
			high = l-1-i
			max = nums[l-1-i]
		}
	}
	if high > low {
		return  high-low+1
	}
	return 0
}

func fmax (x, y int) int {
	if x > y {
		return x
	}
	return y
}

func fmin (x, y int) int {
	if x < y {
		return x
	}
	return y
}
func main() {
	a := []int{9, 8, 4, 8, 10, 9, 15}
	subarray := findUnsortedSubarray(a)
	fmt.Println(subarray)
}

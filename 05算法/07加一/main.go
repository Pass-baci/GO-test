package main

import "github.com/golang/go/src/pkg/fmt"

func plusOne(digits []int) []int {
	var result =  make([]int,1)
	for i := len(digits)-1; i >= 0 ; i-- {
		digits[i] += 1
		if digits[i] == 10 {
			digits[i] = 0
		} else {
			break
		}
	}
	if digits[0] == 0 {
		result[0] = 1
		result = append(result,digits...)
		return result
	} else {
		return digits
	}
}

func main() {
	nums := []int{9,9,8}
	one := plusOne(nums)
	fmt.Println(one)
}

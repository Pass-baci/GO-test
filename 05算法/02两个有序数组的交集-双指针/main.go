package main

import "fmt"

//规则：当指针指向不同的数字时，数字小的指针进行向后移一位
//使用原数组进行编辑交集的数字
func intersect(nums1, nums2 []int) []int {
	i, j, k := 0,0,0 //定义i,j两个指针，k用来指定交集的数字的下标
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			nums1[k] = nums1[i]
			i++
			j++
			k++
		}
	}
	return nums1[:k]
}

func main() {
	num1 := []int{1,3,4,5,6,8,9}
	num2 := []int{1,2,5,6,7,8,9,10,14}
	ints := intersect(num1, num2)
	fmt.Println(ints)
}

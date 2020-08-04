package main


//使用map的Key唯一性
func intersect(nums1, nums2 []int) []int {
	nummap := map[int]int{}
	for _, v := range nums1 {
		nummap[v]++
	}
	k := 0
	for _, v := range nums2 {
		if nummap[v] > 0 {
			nummap[v] -= 1 //因为是交集数，所以减一
			nums2[k] = v
			k++
		}
	}
	return nums2

}
func main() {

}

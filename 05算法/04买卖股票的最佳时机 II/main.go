package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	j := 1
	sum := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] - prices[j] < 0 {
			sum1 := prices[j] - prices[i]
			sum += sum1
		}
		j++
	}
	return sum

}
func main() {
	ars := []int{7,1,5,3,6,4}
	profit := maxProfit(ars)
	fmt.Println(profit)
}

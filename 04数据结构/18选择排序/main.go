package main

import "fmt"

func slectsort(selc *[8]int){
	for i := 0; i < len(selc); i++ {
		max := selc[i]
		maxIndex := i
		for j := i + 1; j< len(selc); j++ {
			if selc[j] < max {
				max = selc[j]
				maxIndex = j
			}
		}
		if maxIndex != i {
			selc[i],selc[maxIndex] = selc[maxIndex],selc[i]
		}
		fmt.Printf("第%d次比较结果为：%v\n",i+1,selc)
	}
}

func main() {
	slec := &[...]int{1,3,5,2,56,6761,323,3434}
	slectsort(slec)
}

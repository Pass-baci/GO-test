package main

import "fmt"

var arrmap [6][6]int

func SetWay(arrmap *[6][6]int, i int, j int) bool {
	if arrmap[4][4] == 2 {
		return true
	} else {
		if arrmap[i][j] == 0 {
			arrmap[i][j] = 2
			if SetWay(arrmap,i+1,j) {
				return true
			} else if SetWay(arrmap,i,j+1){
				return true
			} else if SetWay(arrmap,i-1,j) {
				return true
			} else if SetWay(arrmap,i,j-1) {
				return true
			} else {
				arrmap[i][j] = 3
				return false
			}
		} else {
			return false
		}
	}
}

func main() {
	for i := 0; i < 6; i++ {
		arrmap[0][i] = 1
		arrmap[5][i] = 1
		arrmap[i][0] = 1
		arrmap[i][5] = 1
	}
	SetWay(&arrmap,1,1)
	for i := 0; i < 6; i++ {
		fmt.Println(arrmap[i])
	}
}

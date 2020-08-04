package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type ValNode struct {
	row int
	col int
	val int
}

func main() {
	var chessMap[11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2
	//输出原始数组的样子
	for _, v := range chessMap {
		for _, v1 := range v {
			fmt.Printf("%d\t",v1)
		}
		fmt.Println()
	}
	//原始数组转结构体
	var sparseArr []ValNode //定义一个切片用于储存结构体
	//记录二维数组的规模
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr,valNode)
	for i, v := range chessMap {
		for j, v1 := range v {
			if v1 != 0 {
				valNode := ValNode{
					row: i,
					col: j,
					val: v1,
				}
				sparseArr = append(sparseArr,valNode)
			}
		}
	}
	//将稀疏数组写入到文件中
	file, err := os.OpenFile("E:/Go/src/test_test/12稀疏数组/chessmap.data", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Open file is failed,err:",err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, valNode := range sparseArr {
		sprintf := fmt.Sprintln(valNode.row, valNode.col, valNode.val)
		writer.WriteString(sprintf)
	}
	writer.Flush()
	//恢复成原始数组
	open, err := os.Open("E:/Go/src/test_test/12稀疏数组/chessmap.data")
	if err != nil {
		fmt.Println("Open file is failed err:",err)
		return
	}
	defer open.Close()
	var temp [][]string
	reader := bufio.NewReader(open)
	for {
		readString, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("reader file is failed err:",err)
			return
		}
		replace := strings.Replace(readString,"\n", "", -1)
		split := strings.Split(replace, " ")
		temp = append(temp,split)

	}
	fmt.Println(temp)
	var atoi1 []int
	var chessMap2 [11][11]int
	for i, v := range temp {
		if i != 0 {
			atoi, err := strconv.Atoi(v[i])
			atoi1 = append(atoi1, atoi)
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Println(atoi1)
	}
	for i := 0; i< len(atoi1)-2;i++ {
		chessMap2[atoi1[i]][atoi1[i+1]] = atoi1[i+2]
	}
	fmt.Println(chessMap2)
	fmt.Println(atoi1)
}

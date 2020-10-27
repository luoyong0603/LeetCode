package main

import "fmt"

/*
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
*/

func main() {
	fmt.Print(generate(5))
}

func generate(numRows int) [][]int {
	var l [][]int
	for i := 0; i < numRows; i++ {
		//第0个跟第1个先枚举掉
		if i == 0 {
			var tmp = []int{1}
			l = append(l, tmp)
			continue
		}
		if i == 1 {
			var tmp = []int{1, 1}
			l = append(l, tmp)
			continue
		}
		//当numRows大于2时才执行以下方法
		l = getList(l)
	}
	return l
}

func getList(ll [][]int) [][]int {
	//获取二维数组的最后一个元素 l
	var lastList = ll[len(ll)-1]
	//新一维数组
	var l []int
	for i := 0; i < len(lastList); i++ {
		//第0位跟最后一位，是1
		if i == 0 {
			l = append(l, 1)
			continue
		}
		//中间位是lastList 中的前两数之和
		l = append(l, lastList[i-1]+lastList[i])
		//第0位跟最后一位，是1
		if i == len(lastList)-1 {
			l = append(l, 1)
			continue
		}
	}
	ll = append(ll, l)
	return ll
}

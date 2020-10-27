package main

import "fmt"

/*
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。

注意：你不能在买入股票前卖出股票。

示例 1:

输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
*/

//找到数组中的最小值，与最大值，必须满足最小值在最大值之前

func main() {

	var list = []int{7, 6, 4, 3, 1}
	i := maxProfit(list)
	fmt.Println("i=", i)
}

func maxProfit(prices []int) int {

	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}

	//找最大值
	var max = prices[1] - prices[0]
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			flag := prices[j] - prices[i]
			if max < flag {
				max = flag
			}
		}
	}
	if max < 0 {
		return 0
	}
	return max
}

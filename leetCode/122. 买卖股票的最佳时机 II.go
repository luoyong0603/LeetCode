package main

import "fmt"

func main() {
	var l = []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit2(l))
}

//贪心算法
func maxProfit2(prices []int) int {
	// 设置当前收益为0
	var max = 0
	//假设从第二天开始买入，
	for i := 1; i <= len(prices); i++ {
		//当天价格高于前一天，就卖掉
		if prices[i] > prices[i-1] {
			max += prices[i] - prices[i-1]
		}

	}
	return max
}

package main

import (
	"fmt"
)

func main() {
	var l = []int{1, 1, 1, 1, 2, 2, 2, 3, 3, 3}
	fmt.Print(distributeCandie(l))
}

func distributeCandie(candies []int) int {
	m := make(map[int]int)
	for i := 0; i < len(candies); i++ {
		//目的是计算有多少种不同种类的糖果数（可以看成不同的众数）
		m[candies[i]]++
	}
	//如果众数大于长度的一半，则妹妹可以获得的最大糖果的种类数就是数组长度的一半
	if len(m) > len(candies)/2 {
		return len(candies) / 2
	}
	//如果众数小于map长度的一半，则最大糖果的种类数就是该map长度
	return len(m)
}

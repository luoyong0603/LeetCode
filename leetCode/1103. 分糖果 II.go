package main

import "fmt"

/*

示例 1：

输入：candies = 7, num_people = 4
输出：[1,2,3,1]
解释：
第一次，ans[0] += 1，数组变为 [1,0,0,0]。
第二次，ans[1] += 2，数组变为 [1,2,0,0]。
第三次，ans[2] += 3，数组变为 [1,2,3,0]。
第四次，ans[3] += 1（因为此时只剩下 1 颗糖果），最终数组变为 [1,2,3,1]。


示例 2：

输入：candies = 10, num_people = 3
输出：[5,2,3]
解释：
第一次，ans[0] += 1，数组变为 [1,0,0]。
第二次，ans[1] += 2，数组变为 [1,2,0]。
第三次，ans[2] += 3，数组变为 [1,2,3]。
第四次，ans[0] += 4，最终数组变为 [5,2,3]。
*/

func main() {
	var c = 10
	var k = 3
	fmt.Print(distributeCandies(c, k))
}

func distributeCandies(candies int, num_people int) []int {
	ans := make([]int, num_people)
	//表示当前所需发的糖果数
	count := 0
	//记录遍历次数，方便下标循环
	index := 0
	// 只要糖还有，就继续分
	for candies > 0 {
		count++
		//返回的数会一直是相应数组的下标
		i := index % num_people
		//足够比前一数量加一大
		if candies > count {
			ans[i] += count
		} else {
			//不够，则直接把剩余糖果全分，循环结束
			ans[i] += candies
		}
		//往下分
		index++
		//剩余糖果数
		candies -= count
	}
	return ans
}

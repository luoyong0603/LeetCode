package main

import (
	"fmt"
	"sort"
)

func main() {

	var staple = []int{10, 20, 5}
	var drinks = []int{5, 5, 2}
	var x = 9
	fmt.Print(breakfastNumber2(staple, drinks, x))
}

func breakfastNumber2(staple []int, drinks []int, x int) int {
	const NUMBER = 1000000007
	var count = 0
	sort.Ints(staple)
	sort.Ints(drinks)
	for i := 0; i < len(staple); i++ {
		if staple[i] > x {
			break
		}
		var low = 0
		var high = len(drinks)
		var target = x - staple[i]
		//采用二分查找
		for j := len(drinks) - 1; j >= 0 && low < high; j-- {
			var mid = (low + high) / 2
			//比目标值大，说明目标值在当前mid值左边，接着往左分
			if drinks[mid] > target {
				high = mid
			} else {
				//比目标值小，往右+1查找
				low = mid + 1
			}
		}
		//找到小于等于目标值(target)的元素下标，即[0，low]中的值都满足。
		count += low
	}
	return count % NUMBER
}

func breakfastNumber(staple []int, drinks []int, x int) int {
	const NUMBER = 1000000007
	var count int = 0
	sort.Ints(staple)
	sort.Ints(drinks)
	for i := 0; i < len(staple); i++ {
		if staple[i] > x {
			break
		}
		for j := 0; j < len(drinks); j++ {
			if drinks[j] > x {
				break
			}
			if staple[i]+drinks[j] <= x {
				count++
			}
		}
	}
	return count % NUMBER
}

const NUMBER = 1000000007

func breakfastNumber1(staple []int, drinks []int, x int) int {
	sort.Ints(staple)
	sort.Ints(drinks)

	var count int
for1:
	for i := 0; i < len(staple); i++ {
		for j := 0; j < len(drinks); j++ {
			if staple[i]+drinks[j] <= x {
				count++
			} else {
				if staple[i] > x {
					break for1
				}
				if drinks[j] > x {
					break
				}
			}
		}
	}

	return count % NUMBER
}

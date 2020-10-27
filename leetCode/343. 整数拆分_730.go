package main

import (
	"fmt"
	"time"
)

//给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。

func main() {
	//var list []int

	t1 := time.Now()
	sum := integerBreak(58)
	fmt.Println("耗时=",time.Since(t1), "    sum =", sum)
	//s :=1
	//for _, v := range sum {
	//	fmt.Printf("value=%v,",v)
	//	s = s*v
	//}
	//
	//
	//fmt.Println(s)
}

func integerBreak(n int) int {
	var list []int
	sum1 := returnList(n,0, list)
	s :=1
	for _, v := range sum1 {
		s = s*v
	}
	return s
}

func returnList(n,index int, list []int) []int {

	if n-3 >= 3 {
		index ++
		list = append(list, 3)
		return returnList(n-3,index, list)
	}
	if n==5{
		list = append(list, 2,3)
	}
	if n == 4 {
		list = append(list, 2, 2)
	}
	if index >0{
		if n == 3 {
			list = append(list, 3)
		}
		if n==2 ||n==1 {
			list = append(list, n)
		}
		return list
	}

	if n == 3 {
		list = append(list, 2,1)
	}
	if n == 2 {
		list = append(list, 1,1)
	}
	if n== 1 {
		list = append(list, 1)
	}

	return list

}

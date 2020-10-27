package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("请输入一个数字：")

	var num string
	fmt.Scanln(&num)
	//////读键盘
	// 1,1,2,3,5,8,13,21
	q, _ := strconv.Atoi(num)
	for i := 0; i < q; i++ {
		fmt.Print(digui(i), ",")
	}
}

func digui(x int) int {
	if x <= 1 {
		return 1
	}
	return digui(x-1) + digui(x-2)
}

package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println("牛顿法求平方根:Sqrt(10) = ", Sqrt(10))?
	v := 9.00
	t := 0.01
	value := sqrt1(0, 1, v, t)
	fmt.Printf("暴力破解求得%v开根号为:%v", v, value)

}

func Sqrt(F float64) float64 {
	x := 1.0
	for math.Abs(x*x-F) > 1e-10 {
		x -= (x*x - F) / (2 * x)
	}
	return x
}

//暴力法
/*
参数说明：start：表示从start开始计算，deepth：表示深度，v:目标值， t:误差值
*/
func sqrt1(start float64, depth float64, v float64, t float64) float64 {
	var num float64
	for i := start; i < v; i += depth {
		if i*i <= v {
			num = i
			if depth < t {
				return num
			}
		} else {
			break
		}
	}
	return sqrt1(num, depth/10, v, t)
}

//二分法
func sqrt2(v int, t float64) float64 {
	vFloat := float64(v)
	var left, right float64 = 0, vFloat
	for left < right {
		mid := (left + right) / 2
		if mid*mid <= vFloat {
			left = mid
		} else {
			right = mid
		}
		if math.Abs(vFloat-mid*mid) <= t {
			return mid
		}
	}
	return 0
}

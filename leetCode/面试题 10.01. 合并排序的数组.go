package main

import (
	"fmt"
	"sort"
	time "time"
)

func main() {
	var a = []int{1, 2, 3,4,5,6,3,4}
	var b = []int{2, 5, 6,2,3,6,3,9}
	var m = len(a)
	var n = len(b)
	time1:=time.Now()
	time2:=time.Now()
	time3:=time.Now()
	i := merge(a, m, b, n)
	for _, v := range i {
		fmt.Print(v)
	}
	fmt.Println("t1=",time.Since(time1))
	j:=merge1(a, m, b, n)
	for _, v := range j {
		fmt.Print(v)
	}
	fmt.Println("t2=",time.Since(time2))
	k:=merge2(a, m, b, n)
	for _, v := range k {
		fmt.Print(v)
	}
	fmt.Println("t3=",time.Since(time3))
	


}

//选择插入法  双指针
func merge(A []int, m int, B []int, n int) []int {
	var list []int

	//for i,j:=0,0;i<m ||j<n;{
	for i := 0; i < len(A); {
		for j := 0; j < len(B); {
			if i == m {
				list = append(list, B[j])
				j++
			} else if j == n {
				list = append(list, A[i])
				i++
			} else if A[i] > B[j] {
				list = append(list, B[j])
				j++
			} else {
				list = append(list, A[i])
				i++
			}
		}
	}
	//copy(A, list)
	return A
}

//合并后排序法
func merge1(A []int, m int, B []int, n int) []int {
	//var list []int
	A = append(A, B...) //将[]b拼接到[]a后边
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if A[i] > A[j] {
				var temp = A[i]
				A[i] = A[j]
				A[j] = temp
			}
		}
	}
	return A
}

//走捷径  原理还是合并AB再排序

func merge2(A []int, m int, B []int, n int) []int {
	copy(A[m:], B[:])
	sort.Ints(A)
	return A
}

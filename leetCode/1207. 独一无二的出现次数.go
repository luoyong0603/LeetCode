package main

import (
	"fmt"
)

func main() {

	var l = []int{-5, -6, 2, 6, -5, -7, 5}
	fmt.Print(uniqueOccurrences2(l))
}

//map统计出现次数
//map转list
//list去重
func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}
	l := make([]int, 0, len(m))
	for _, v := range m {
		l = append(l, v)
	}
	if len(l) == 1 {
		return true
	}
	for i := 0; i < len(l); i++ {
		for j := i + 1; j < len(l); j++ {
			if l[i] == l[j] {
				return false
			}
		}
	}
	return true
}

//第一个map统计出现次数
//第二个map查重
func uniqueOccurrences1(arr []int) bool {
	m := make(map[int]int)
	m1 := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}
	for _, v := range m {
		_, ok := m1[v]
		if ok {
			return false
		}
		m1[v]++
	}
	return true
}

//第一个list统计出现次数,
func uniqueOccurrences2(arr []int) bool {
	l := make([]int, 2001)
	for _, v := range arr {
		//目的是从范围值开始，最小-1000,
		l[v+1000]++
	}
	for i := 0; i < len(l); i++ {
		if l[i] == 0 {
			continue
		}
		for j := i + 1; j < len(l); j++ {
			if l[i] == 0 {
				continue
			}
			if l[i] == l[j] {
				return false
			}
		}
	}
	return true
}

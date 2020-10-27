package main

import (
	"fmt"
)

/*
给定仅有小写字母组成的字符串数组 A，返回列表中的每个字符串中都显示的全部字符（包括重复字符）组成的列表。
例如，如果一个字符在每个字符串中出现 3 次，但不是 4 次，则需要在最终答案中包含该字符 3 次。
你可以按任意顺序返回答案。
*/

func main() {
	var l = []string{"acabcddd", "bcbdbcbd", "baddbadb", "cbdddcac", "aacbcccd", "ccccddda", "cababaab", "addcaccd"}
	fmt.Print(commonChars(l))
}

func commonChars(A []string) []string {
	m := make(map[string]int)
	var l []string
	//用于初始化m1
	var index = 0
	for i := 0; i < len(A); i++ {
		index++
		m = getOther(index, m, A[i])
	}
	//m 即返回的 公共字符个数
	for k, v := range m {
		//大于1的既多append几次
		if v > 1 {
			for i := 0; i < v; i++ {
				l = append(l, k)
			}
		} else {
			l = append(l, k)
		}
	}
	return l
}

func getOther(index int, m map[string]int, s string) map[string]int {
	//用于统计s字符串的字母数量集
	m1 := make(map[string]int)
	//m2表示两字符串交集的结果集
	m2 := make(map[string]int)
	for i := 0; i < len(s); i++ {
		m1[string(s[i])]++
	}
	//初始化，一开始交集为空，所以用index作为条件判断，把第一个字符串的字母数量集返回过去
	if index == 1 {
		return m1
	}
	for k, v1 := range m1 {
		//判断新字符串中是否含有交集中的字母
		v2, ok := m[k]
		//如果存在，则比较所带的字母个数，取交集（个数少即公共元素）
		if ok {
			if v1 <= v2 {
				m2[k] = v1
				continue
			}
			m2[k] = v2
		}
	}
	//返回交集
	return m2
}

package main

import "fmt"

func main() {

	s := "aaaaaa"
	t := ""
	bool := isSubsequence(s, t)
	fmt.Println(bool)

}

/*
	思路：以短串为目标，遍历长串，每次只需拿短串第0个去匹配长串，如果匹配成功，则以当前下标截取，
          以截取后的数组递归传入，继续
*/
func isSubsequence(s string, t string) bool {
	//用于标记是否找到短串最后一个元素
	var temp = false

	if s == "" {
		return true
	}

	for j := 0; j < len(t); j++ {
		if s[0] == t[j] {
			//如果短串长度为1了，表示已经找到最后一个短串元素
			if len(s) == 1 {
				temp = true
				break
			}
			return isSubsequence(s[1:len(s)], t[j+1:len(t)])
		}
	}
	if temp {
		return true
	}
	return false
}

//双指针解法

func isSubsequence1(s string, t string) bool {
	if s == "" {
		return true
	}
	if t == ""{
		return true
	}
	var i, j int
	for i = 0; i < len(s); {

		if s[i] == t[j] {
			if len(s) == i+1 {
				return true
			}
			i++
		}
		if len(t) == j+1 {
			return false
		}
		j++
	}
	return false

}

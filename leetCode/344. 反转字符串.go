package main

import "fmt"

func main() {
	var l = []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Print(l)
	reverseString1(l)
	fmt.Print(l)
}

//直接暴力破解
func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

//双子针法
func reverseString1(s []byte) {
	for left, right := 0, len(s)-1; left < right; left++ {
		s[left], s[right] = s[right], s[left]
		right--
	}

}

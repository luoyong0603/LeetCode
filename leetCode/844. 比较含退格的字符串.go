package main

import (
	"fmt"
)

func main() {

	s1 := "ab#c"
	s2 := "ad#c"
	fmt.Print(backspaceCompare(s1, s2))
}

func backspaceCompare(S string, T string) bool {
	return Com(S) == Com(T)
}

func Com(s string) string {
	var stack = make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] != '#' {
			stack = append(stack, s[i])
		} else if len(stack) != 0 {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}

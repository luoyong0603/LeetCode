package main

import "fmt"

func main() {

	circle := judgeCircle("LRL")
	fmt.Print(circle)
}

func judgeCircle(moves string) bool {
	var a, b int
	for i := 0; i < len(moves); i++ {
		if moves[i] == 'L' {
			a--
		}
		if moves[i] == 'R' {
			a++
		}
		if moves[i] == 'U' {
			b++
		}
		if moves[i] == 'D' {
			b--
		}
	}
	return a == 0 && b == 0
}

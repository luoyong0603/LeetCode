package main

func calculate(s string) int {
	var x, y = 1, 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "A" {
			x = 2*x + y
		}
		if string(s[i]) == "B" {
			y = 2*y + x
		}
	}
	return x + y
}

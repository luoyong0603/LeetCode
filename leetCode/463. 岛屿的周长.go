package main

import "fmt"

func main() {
	//var list = [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}
	var list1 = [][]int{{1}}
	fmt.Print(islandPerimeter(list1))
}


/**
解题思路，逢1加4，然后再去减临边；就是上下为1，或者左右为1的情况，就减2，
 */
func islandPerimeter(grid [][]int) int {
	var sum int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				sum += 4
			}
			//过滤二维数组只有一个一位数组的情况
			if i >= 1 {
				if grid[i][j] == 1 && grid[i-1][j] == 1 {
					sum -= 2
				}
			}
			//过滤二维数组中一位数组中只有一位的情况
			if j >= 1 {
				if grid[i][j] == 1 && grid[i][j-1] == 1 {
					sum -= 2
				}
			}
		}
	}
	return sum
}

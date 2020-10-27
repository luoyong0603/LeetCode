package main

import "fmt"

/**
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。
示例:
输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。
*/
func main() {
	var isArr = [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}

	pathSum := minPathSum(isArr)
	fmt.Println(pathSum)

}
func minPathSum(grid [][]int) int {
	x := len(grid)
	y := len(grid[0])
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				grid[0][j] += grid[0][j-1]
			} else if j == 0 {
				grid[i][0] += grid[i-1][0]
			} else {
				grid[i][j] += ReturnMin(grid[i-1][j], grid[i][j-1])
			}

		}
	}
	return grid[x-1][y-1]

}

func ReturnMin(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

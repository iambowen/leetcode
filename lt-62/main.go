package main

import "fmt"

func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		grid[0][i] = 1
	}
	for i := 0; i < m; i++ {
		grid[i][0] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}
	return grid[m-1][n-1]
}
func main() {
	fmt.Println(uniquePaths(3, 2), uniquePaths(7, 3))
}

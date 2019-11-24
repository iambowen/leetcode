package main

import "fmt"

const INF = 2147483647

func wallsAndGates(rooms [][]int) {
	if len(rooms) == 0 {
		return
	}
	m, n := len(rooms), len(rooms[0])
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	queue := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rooms[i][j] == 0 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	for len(queue) != 0 {
		temp := queue
		queue = [][]int{}
		for i := 0; i < len(temp); i++ {
			point := temp[i]
			row, col := point[0], point[1]

			for i := 0; i < len(directions); i++ {
				r := row + directions[i][0]
				c := col + directions[i][1]
				if r < 0 || r >= m || c < 0 || c >= n || rooms[r][c] != INF {
					continue
				}
				rooms[r][c] = rooms[row][col] + 1
				queue = append(queue, []int{r, c})
			}
		}
	}
}

func main() {
	directions := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	fmt.Println(len(directions))
}

package dfs

func maxAreaOfIsland(grid [][]int) int {
	result := 0
	for i := range grid {
		for j := range grid[i] {
			area := checkLand(grid, i, j)
			if area > result {
				result = area
			}
		}
	}
	return result
}

func checkLand(grid [][]int, x, y int) int {
	if grid[x][y] == 0 {
		return 0 // already searched
	}

	grid[x][y] = 0
	area := 1
	// left, right up ,down searched
	if x > 0 {
		area += checkLand(grid, x-1, y)
	}
	if x < len(grid)-1 {
		area += checkLand(grid, x+1, y)
	}

	if y > 0 {
		area = checkLand(grid, x, y-1)
	}
	if y < len(grid)-1 {
		area += checkLand(grid, x, y+1)
	}
	return area
}

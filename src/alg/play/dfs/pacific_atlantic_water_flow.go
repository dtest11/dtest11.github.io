package dfs

/**

Given the following 5x5 matrix:

  Pacific ~   ~   ~   ~   ~
       ~  1   2   2   3  (5) *
       ~  3   2   3  (4) (4) *
       ~  2   4  (5)  3   1  *
       ~ (6) (7)  1   4   5  *
       ~ (5)  1   1   2   4  *
          *   *   *   *   * Atlantic

Return:

[[0, 4], [1, 3], [1, 4], [2, 2], [3, 0], [3, 1], [4, 0]]
 (positions with parentheses in above matrix).

虽然题目要求的是满足向下流能到达两个大洋的位置,如果我们对所有的位置进行搜索,那
么在不剪枝的情况下复杂度会很高。因此我们可以反过来想,从两个大洋开始向上流,这样我们
只需要对矩形四条边进行搜索。搜索完成后,只需遍历一遍矩阵,满足条件的位置即为两个大洋
向上流都能到达的位置。

https://www.youtube.com/watch?v=blOc3KlAR2I
https://leetcode.com/problems/pacific-atlantic-water-flow/



X,Y轴结合分析


**/

func pacificAtlantic(matrix [][]int) [][]int {
	var result [][]int
	if matrix == nil || len(matrix[0]) == 0 {
		return result
	}
	colN := len(matrix[0])
	rowN := len(matrix)
	pacific := make([][]bool, rowN)
	for i := 0; i < rowN; i++ {
		pacific[i] = make([]bool, rowN)
	}
	atlantic := make([][]bool, rowN)
	for i := 0; i < rowN; i++ {
		atlantic[i] = make([]bool, colN)
	}
	// top 	and bottom
	for col := 0; col < colN-1; col++ {
		pacificAtlantaDFS(matrix, 0, col, matrix[0][col], pacific)          //top
		pacificAtlantaDFS(matrix, rowN-1, col, matrix[0][rowN-1], atlantic) // bottom
	}

	// left And right
	for row := 0; row < rowN-1; row++ { //行
		pacificAtlantaDFS(matrix, row, 0, matrix[row][0], pacific)            //left
		pacificAtlantaDFS(matrix, row, colN-1, matrix[row][colN-1], atlantic) // right
	}
	for i := 0; i < rowN; i++ {
		for j := 0; j < colN; j++ {
			if pacific[i][j] && atlantic[i][j] {
				tmp := [][]int{{i, j}}
				result = append(result, tmp...)
			}
		}
	}
	return result
}

// col 列， row行
func pacificAtlantaDFS(matrix [][]int, row int, col int, preHeight int, ocean [][]bool) {
	if row > len(matrix)-1 ||
		row < 0 ||
		col > len(matrix[0])-1 ||
		col < 0 ||
		matrix[row][col] < preHeight ||
		ocean[row][col] {
		return
	}
	ocean[row][col] = true
	pacificAtlantaDFS(matrix, row+1, col, matrix[row][col], ocean)
	pacificAtlantaDFS(matrix, row-1, col, matrix[row][col], ocean)
	pacificAtlantaDFS(matrix, row, col+1, matrix[row][col], ocean)
	pacificAtlantaDFS(matrix, row, col-1, matrix[row][col], ocean)
}

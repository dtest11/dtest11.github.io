package dfs

import (
	"testing"
)

func Test_pacificAtlantic(t *testing.T) {
	m := 10
	n := 2 //[10][2]// 10行2列
	pacific := make([][]bool, m)
	for i := 0; i < m; i++ {
		pacific[i] = make([]bool, n)
	}
	t.Logf("%+v", pacific)
	t.Logf("%+v", [10]int{})
}

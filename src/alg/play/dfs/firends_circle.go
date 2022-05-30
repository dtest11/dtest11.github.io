package dfs

//func findCircleNum(M [][]int) int {
//	if M == nil {
//		return 0
//	}
//	row := len(M)
//	ans := 0
//	for i := 0; i < row; i++ {
//		if M[i][i] != 0 {
//			continue
//		}
//		ans++
//		dfs(M, i, n)
//	}
//
//}
//
//func dfs(m [][]int, cur, n int) {
//	for i := 0; i < n; i++ {
//		if m[i][i] != 0 {
//			continue
//		}
//		//m[cur][i]=m[i][cur]=0
//		dfs(m, i, n)
//	}
//}

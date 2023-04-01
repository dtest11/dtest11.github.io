package twopointers

/*
*
从l位置，开始 ，可以到达的距离， 选择某一个距离，继续往下走，使用同样的贪心策略
i + minJump <= j <= min(i + maxJump, s.length - 1), and

*
*/
func canReach(s string, minJump int, maxJump int) bool {
	//var cur = 0
	//n := len(s)
	//for cur < len(s)-1 {
	//	val := s[cur]
	//	i, _ := strconv.Atoi(string(val))
	//
	//	l := i + minJump
	//	r := i + maxJump
	//
	//	if r < len(s)-1 {
	//		r = len(s) - 1
	//	}
	//	// 此时可以选择 将cur 移动到l..r中的任意一个位置
	//	// 贪心策略每次移动最多，重复，如果不能实现，那么移动r -1 个位置，r-2个位置
	//
	//}
	panic("not implement")
}

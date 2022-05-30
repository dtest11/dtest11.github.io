package union_set

import (
	"strings"
)

// 判断合法算式
/**
给你一个数组equations，装着若干字符串表示的算式。每个算式equations[i]长度都是 4，而且只有这两种情况：a==b或者a!=b，其中a,b可以是任意小写字母。你写一个算法，如果equations中所有算式都不会互相冲突，返回 true，否则返回 false。

比如说，输入["a==b","b!=c","c==a"]，算法返回 false，因为这三个算式不可能同时正确。

再比如，输入["c==c","b==d","x!=z"]，算法返回 true，因为这三个算式并不会造成逻辑冲突。

我们前文说过，动态连通性其实就是一种等价关系，具有「自反性」「传递性」和「对称性」，其实==关系也是一种等价关系，具有这些性质。所以这个问题用 Union-Find 算法就很自然。
*/

func equationsPossible(data []string) bool {
	var unionFind = initUnion(26)
	// 先获取== 建立并查集
	for _, v := range data {
		index := strings.Index(v, "==")
		if index != -1 {
			first := v[0] - 'a'
			last := v[3] - 'a'
			unionFind.union(int(first), int(last))
		}
	}

	// 检查
	for _, v := range data {
		index := strings.Index(v, "!=")
		if index != -1 {
			first := v[0] - 'a'
			last := v[3] - 'a'
			if unionFind.connected(int(first), int(last)) {
				return false
			}
		}
	}
	return true
}

func strToInt(str string) (sum int) {
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		sum += int(runes[i])
	}
	return sum
}

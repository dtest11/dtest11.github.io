package slidewindow


func removeDumplicateEle(input string) string{
	// sort input by assica
	slow :=0
	fast :=0
	for :fast<len(input);fast++ {
		for input[slow] !=input[fast]{
			input[slow]=input[fast]
			slow ++
		}
	}
	return input[:slow+1]
}

func removeTarget(input string,target string) string{
	// 核心思路还是双指针的 快慢，如果fast==target 那么
	// slow 不增长,一直移动直到 找到一个合法的直，将这个值移动到slow的index上
	slow :=0
	for fast:=0;fast <len(input);fast++{
		if input[fast] !=target{
			input[slow]=input[slow]
			slow++
		}
	}
	return input[:slow+1]
}



// 删除重复的元素：相同元素最到只能保留2个
func removeThirdDuplicateEle(input []string) string {
	slow :=0 
	for fast:=0;fast<len(input);fast++ {
		// 比较prev=cur<next
		if input[fast-1]==input[fast] && input[fast]== input[fast+1]{
			continue
		}
		input[slow]=input[fast]
		slow++
	}
}

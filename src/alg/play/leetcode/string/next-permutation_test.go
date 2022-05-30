package string

/**
题解
如 1.2.3
这个排列的组合的结果:
1,2,3
1,3,2
2,1,3
2,3,1
3,1,2
3,2,1

如上面这个组合：最后一个结果是降序（3，2，1）
这样，排列 [2,3,1][2,3,1] 的下一个排列即为 [3,1,2][3,1,2]。特别的，
最大的排列 [3,2,1][3,2,1] 的下一个排列为最小的排列 [1,2,3][1,2,3]。

注意到下一个排列总是比当前排列要大，除非该排列已经是最大的排列。
我们希望找到一种方法，能够找到一个大于当前序列的新序列，
且变大的幅度尽可能小。具体地：
1.
我们需要将一个左边的「较小数」与一个右边的「较大数」交换，
以能够让当前排列变大，从而得到下一个排列。
2.
同时我们要让这个「较小数」尽量靠右，而「较大数」尽可能小。
当交换完成后，「较大数」右边的数需要按照升序重新排列。
这样可以在保证新排列大于原来排列的情况下，使变大的幅度尽可能小。


以排列 [4,5,2,6,3,1][4,5,2,6,3,1] 为例：
我们能找到的符合条件的一对「较小数」与「较大数」的组合为 22 与 33，满足「较小数」尽量靠右，而「较大数」尽可能小。
当我们完成交换后排列变为 [4,5,3,6,2,1][4,5,3,6,2,1]，此时我们可以重排「较小数」右边的序列，序列变为
[4,5,3,1,2,6][4,5,3,1,2,6]。

具体地，我们这样描述该算法，对于长度为 nn 的排列 aa：
1.
首先从后向前查找第一个顺序对 (i,i+1)(i,i+1)，满足 a[i] < a[i+1]
这样「较小数」即为 a[i]。此时 [i+1,n) 必然是下降序列

2.
如果找到了顺序对，那么在区间 [i+1,n)中从后向前查找第一个元素 j 满足
a[i] < a[j]。这样「较大数」即为 a[j]。
3.
交换 a[i]a[i] 与 a[j]a[j]，此时可以证明区间 [i+1,n) 必为降序。
我们可以直接使用双指针反转区间 [i+1,n) 使其变为升序，而无需对该区间进行排序。



*/

func nextPermutation(nums []int) {
	length := len(nums)
	smallIndex := length - 2
	for smallIndex >= 0 && nums[smallIndex] >= nums[smallIndex+1] {
		smallIndex--
	}
	if smallIndex >= 0 {
		j := smallIndex - 1
		for j >= 0 && nums[j] <= nums[smallIndex] {
			j--
		}
		//
		nums[j], nums[smallIndex] = nums[smallIndex], nums[j]
	}
	// reverse smallIndex-j
	reverse(nums[smallIndex+1:])
}

func reverse(a []int) {
	right := len(a) - 1
	for i := 0; i < len(a)/2; i++ {
		a[i], a[right] = a[right], a[i]
		right--
	}
}

/**
31. Next Permutation
Medium



5946

1992

Add to List

Share
Implement next permutation, which rearranges numbers into the
lexicographically next greater permutation of numbers.

If such an arrangement is not possible, it must rearrange it as the lowest
 possible order (i.e., sorted in ascending order).

The replacement must be in place and use only constant extra memory.



Example 1:

Input: nums = [1,2,3]
Output: [1,3,2]
Example 2:

Input: nums = [3,2,1]
Output: [1,2,3]
Example 3:

Input: nums = [1,1,5]
Output: [1,5,1]
Example 4:

Input: nums = [1]
Output: [1]

**/

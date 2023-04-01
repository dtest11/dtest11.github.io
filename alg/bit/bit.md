```text
符号	描述	运算规则
&	与	两个位都为1时，结果才为1
|	或	两个位都为0时，结果才为0
^	异或	两个位相同为0，相异为1
~	取反	0变1，1变0
<<	左移	各二进位全部左移若干位，高位丢弃，低位补0
>>	右移	各二进位全部右移若干位，对无符号数，高位补0，有符号数，各编译器处理方法不一样，有的补符号位（算术右移），有的补0（逻辑右移）

```
## [136. Single Number](https://leetcode.com/problems/single-number/description/)

给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。

异或运算的法则类似乘法：满足交换律、结合律。
```text
a ^ b = b ^ a (a ^ b) ^ c = a ^ (b ^ c)
```

```go 
func singleNumber(nums []int) int {
    var res int
    for _,v :=range nums{
       res = res^v
    }
   return res
}
```
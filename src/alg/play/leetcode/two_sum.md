## Two sum 类型

1. 主要的解决的方法就是 滑动窗口, sorted 
```text

left :=0
right :=len(array)-1

for left !=right {
     start :=array[left]
     end := array[right]
     sum :=start+end
     if sum > k{
         right --
     }
     if sum == k {
         return true
     }
     if sum <k{
         left++
     }
}


2. hashSet  not sorted

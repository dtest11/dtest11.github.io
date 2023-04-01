滑动窗口
```go
func slide(s string) string {
    left :=0
    right :=0
    window :=[]string
    for right < s.len(){
      window.append(s[right])
      right++
      for 滑动条件成立 {
          windows.remove(s[right])
          left++
      }
    }
}
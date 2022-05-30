package main

import (
	"fmt"
	"testing"
)

func TestZslRandomLevel(t *testing.T) {
	// for i := 0; i < 10000; i++ {
	// 	t.Log(ZslRandomLevel())
	// 	level := ZslRandomLevel()
	// 	assert.Equal(t, true, level > 0, "The two words should be the same.")
	// }
	abc := make(map[string]string)
	fmt.Println(abc["name"], "-------")
}

func TestZslInsert(t *testing.T) {
	sl := zslCreate()
	ZslInsert(sl, 2, "2")
	fmt.Println("---插入的数据 2 ： 2")
	read(sl)
	fmt.Println("---插入的数据 5 ： 5")
	ZslInsert(sl, 5, "5")
	read(sl)

	fmt.Println("---插入的数据 7 ： 7")
	ZslInsert(sl, 7, "7")
	read(sl)

	fmt.Println("---插入的数据 8 ： 8")
	ZslInsert(sl, 9, "9")
	read(sl)

	fmt.Println("---插入的数据 9 ： 9")
	ZslInsert(sl, 12, "12")
	read(sl)
}

func read(sl *zskiplist) {
	fmt.Println("********打印结果*********")
	fmt.Println("                     ")
	fmt.Println("                     ")
	h := sl.Header
	MaxLevel := sl.Level
	for h != nil && MaxLevel >= 0 {
		fmt.Printf("----- 遍历第 %d 层-------\n", MaxLevel+1)
		// 当前层向前推进
		if h.Level != nil && h.Level[MaxLevel] != nil && h.Level[MaxLevel].Forward != nil {
			fmt.Print("该层的数据有-----  ele= ", h.Level[MaxLevel].Forward.Ele, ":score=", h.Level[MaxLevel].Forward.Score, "   span=", h.Level[MaxLevel].Span)
			h = h.Level[MaxLevel].Forward
			fmt.Println(" 继续向前推进")
			continue
		}
		// 移动到下一层
		MaxLevel = MaxLevel - 1
	}
	fmt.Println("                     ")
	fmt.Println("                     ")
}

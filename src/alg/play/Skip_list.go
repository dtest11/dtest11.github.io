package main

import (
	"log"
	"math/rand"
	"strconv"
	"time"
)

/***
Redis的skiplist第一層是一個雙向列表 原文網址：https://itw01.com/UYYNYER.html
redis 产生的level 永远不会是0

*/

func sdscmp(a, b string) int {
	if a == b {
		return 0
	}
	aI, _ := strconv.ParseInt(a, 10, 64)
	bI, _ := strconv.ParseInt(b, 10, 64)
	log.Println(aI-bI)
	return int(aI - bI)
}


type zskiplistLevel struct {
	Forward *zskiplistNode
	Span    int
}
type zskiplistNode struct {
	Ele      string
	Score    int
	BackWard *zskiplistNode //向后
	Level    []*zskiplistLevel
}

type zskiplist struct {
	Header *zskiplistNode
	Tail   *zskiplistNode
	Length int
	Level  int
}

const (
	ZSKIPLIST_MAXLEVEL = 4
	ZSKIPLIST_P        = 0.25
)

func zslCreateNode(level int, score int, ele string) *zskiplistNode {
	zslNode := zskiplistNode{
		Ele:      ele,
		Score:    score,
		BackWard: nil,
		Level:    make([]*zskiplistLevel, level),
	}
	for v := range zslNode.Level {
		zslNode.Level[v] = &zskiplistLevel{
			Forward: nil,
			Span:    0,
		}
	}
	return &zslNode
}

/* Create a new skiplist. */
func zslCreate() *zskiplist {
	//var j int
	var zsl zskiplist
	zsl.Level = 1
	zsl.Length = 0
	zsl.Header = zslCreateNode(ZSKIPLIST_MAXLEVEL, 0, "")
	////每一层都设置下
	//for ; j < ZSKIPLIST_MAXLEVEL; j++ {
	//	zsl.Header.Level[j].Forward = nil
	//	zsl.Header.Level[j].Span = 0
	//}
	//zsl.Header.BackWard = nil
	//zsl.Tail = nil

	return &zsl
}

/* Returns a random level for the new skiplist node we are going to create.
 * The return value of this function is between 1 and ZSKIPLIST_MAXLEVEL
 * (both inclusive), with a powerlaw-alike distribution where higher
 * levels are less likely to be returned. */
func ZslRandomLevel() int {
	level := 1
	val := 0xFFFF
	rand.Seed(time.Now().UnixNano())

	for rand.Int()&val < val/4 {
		level += 1
	}
	if level >= ZSKIPLIST_MAXLEVEL {
		level = ZSKIPLIST_MAXLEVEL - 1
	}
	log.Println("------新产生的level----:  ", level)

	return level
}

/* Insert a new node in the skiplist. Assumes the element does not already
 * exist (up to the caller to enforce that). The skiplist takes ownership
 * of the passed SDS string 'ele'. */
func ZslInsert(zsl *zskiplist, score int, ele string) *zskiplist {
	var (
		update [ZSKIPLIST_MAXLEVEL]*zskiplistNode // 保存搜素路径
		x      *zskiplistNode
		rank   [ZSKIPLIST_MAXLEVEL]int

		i     int
		level int
	)

	x = zsl.Header
	// 从最高层开始循环，
	for i = zsl.Level - 1; i >= 0; i-- {
		if i == zsl.Level-1 { // i+1=zsl.Level ,那么i = 第 最高的一层
			rank[i] = 0
		} else {
			rank[i] = rank[i+1] //
		}

		// 比较key 对应的value
		// 沿着前进指针遍历跳跃表， 对比分值，
		if (x.Level[i] != nil && x.Level[i].Forward != nil) && (x.Level[i].Forward.Score <= score && sdscmp(x.Level[i].Forward.Ele, ele) < 0) {
			rank[i] += x.Level[i].Span // 记录中途跨越了多少个节点
			x = x.Level[i].Forward     // 移动至下一个指针
		}
		update[i] = x // 记录将要和新节点相连接的节点， 记录每一层的数据， 在查询的过程中，每一层最后碰到的数据
	}
	level = ZslRandomLevel()
	// 如果新节点的层数比表中其他节点的层数都要大 ，那么初始化表头节点中未使用的层，并将它们记录到 update 数组中 将来也指向新节点
	if level > zsl.Level {
		// 初始化未使用层
		// T = O(1)
		for i = zsl.Level; i < level; i++ {
			rank[i] = 0
			update[i] = zsl.Header // 没有数据，初始化为header
			update[i].Level[i].Span = zsl.Length
		}
		zsl.Level = level // 更新表中节点最大层数
	}

	x = zslCreateNode(level, score, ele) //创建新节点
	// 将前面记录的指针指向 新节点 并做相应的设置
	for i = 0; i < level; i++ {
		// 标准的 链表中插入一个元素的操作
		x.Level[i].Forward = update[i].Level[i].Forward // 设置新节点的 forward 指针
		update[i].Level[i].Forward = x                  // 将沿途记录的各个节点的 forward 指针指向新节点

		/* update span covered by update[i] as x is inserted here */
		x.Level[i].Span = update[i].Level[i].Span - (rank[0] - rank[i]) // 计算新节点跨越的节点数量
		update[i].Level[i].Span = (rank[0] - rank[i]) + 1               // 更新新节点插入之后，沿途节点的 span 值
		// 其中的 +1 计算的是新节点
	}

	/* increment span for untouched levels */
	for i = level; i < zsl.Level; i++ {
		update[i].Level[i].Span++ // 增加span 由于是新产生的节点

	}
	//设置新节点的 后退指针， 第一层是双向链表
	if update[0] == zsl.Header {
		x.BackWard = nil
	} else {
		x.BackWard = update[0]
	}
	if x.Level != nil && x.Level[0] != nil && x.Level[0].Forward != nil {
		x.Level[0].Forward.BackWard = x
	}

	// 跳跃表的节点计数增一
	zsl.Length++
	return zsl
}

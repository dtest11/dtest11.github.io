package v1

import "container/list"

type LFUCache struct {
	capacity  int
	freqTable map[int]*list.List
	keyTable  map[int]*list.Element
	miniFreq  int // 最小的频率 当前
}

type Payload struct {
	key  int
	val  int
	freq int
}

func Constructor(capacity int) LFUCache {
	l := LFUCache{
		capacity:  capacity,
		freqTable: make(map[int]*list.List, capacity),
		keyTable:  make(map[int]*list.Element, capacity),
	}
	return l
}

// increaseFreq 添加ele 对应的频率，同时更新miniFrequency
func (lfu *LFUCache) increaseFreq(ele *list.Element, payload Payload) {
	if lfu.miniFreq == payload.freq {
		// 当前的该频率的元素只有一个，那么如果移动了频率，最小频率也需要改变，不然找不到
		if lfu.freqTable[payload.freq].Len() == 1 {
			lfu.miniFreq = lfu.miniFreq + 1
		}
	}
	// 1. 删除频率对应的table
	lfu.freqTable[payload.freq].Remove(ele)
	// 2.移动该元素到freq+1 对应的table
	payload.freq = payload.freq + 1
	// 2.1 检查table是否没有初始化
	if _, ok := lfu.freqTable[payload.freq]; !ok {
		lfu.freqTable[payload.freq] = list.New()
	}
	ele = lfu.freqTable[payload.freq].PushBack(payload)
	//  3. 更新key 对应的table
	lfu.keyTable[payload.key] = ele
}

func (lfu *LFUCache) RemoveMini() {
	table := lfu.freqTable[lfu.miniFreq]
	front := table.Front()
	table.Remove(front)
	delete(lfu.keyTable, front.Value.(Payload).key)
}

func (lfu *LFUCache) Get(key int) int {
	if ele, ok := lfu.keyTable[key]; ok {
		payload := ele.Value.(Payload)
		lfu.increaseFreq(ele, payload)
		return payload.val
	}
	return -1
}

func (lfu *LFUCache) Put(key int, value int) {
	if ele, ok := lfu.keyTable[key]; ok {
		payload := ele.Value.(Payload)
		payload.val = value
		lfu.increaseFreq(ele, payload)
		return
	}
	if lfu.capacity == len(lfu.keyTable) {
		lfu.RemoveMini()
	}
	lfu.miniFreq = 1
	payload := Payload{
		key:  key,
		val:  value,
		freq: 1,
	}
	var l *list.List
	l, ok := lfu.freqTable[1]
	if !ok {
		l = list.New()
		lfu.freqTable[1] = l
	}
	ele := l.PushBack(payload)
	lfu.keyTable[key] = ele
}

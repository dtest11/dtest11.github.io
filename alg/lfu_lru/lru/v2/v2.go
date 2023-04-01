package v2

type LRUCache struct {
	capacity int
	keyTable map[int]*Element  // get 操作保证1的时间复杂度
	list     *DoublyLinkedList // 保证数据的顺序，最近使用的在front ,最后使用的在end
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		keyTable: make(map[int]*Element),
		list:     NewDoublyLinkedList(),
	}
}

func (lru *LRUCache) Get(key int) int {
	if ele, ok := lru.keyTable[key]; ok {
		// 移动到最前面的，最前面的元素是最近使用的
		payload := ele.Value.(Payload)
		lru.moveToFront(ele, payload)
		return payload.val
	}
	return -1
}

func (lru *LRUCache) Put(key int, value int) {
	if ele, ok := lru.keyTable[key]; ok {
		payload := ele.Value.(Payload)
		payload.val = value
		lru.moveToFront(ele, payload)
		return
	}
	if lru.capacity == len(lru.keyTable) {
		// delete 最后一个元素
		tail := lru.list.Back()
		lru.list.Remove(tail)
		delete(lru.keyTable, tail.Value.(Payload).key) // 注意删除链表和字典
	}
	// 元素插入到头部
	payload := Payload{
		key: key,
		val: value,
	}
	ele := lru.list.PushFront(payload)
	lru.keyTable[payload.key] = ele
}

// moveKey 移动到最前面的，最前面的元素是最近使用的
func (lru *LRUCache) moveToFront(ele *Element, payload Payload) {
	lru.list.Remove(ele)
	ele = lru.list.PushFront(payload)
	lru.keyTable[payload.key] = ele
}

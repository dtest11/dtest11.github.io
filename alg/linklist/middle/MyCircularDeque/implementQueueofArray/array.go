package mycirculardeque

/*
*使用数组 模拟循环队列/
 */
type MyCircularDeque struct {
	size  int
	queue *queue
}

type queue struct {
	size  int
	array []int
	front int // 指向最早插入的元素
	rear  int // 指向最后一个元素的下一个元素
}

func (q *queue) IsEmpty() bool {
	return q.front == q.rear
}

func (q *queue) IsFull() bool {
	return (q.rear+1)%q.size == q.front
}

func (q *queue) enqueue(val int) bool {
	if q.IsFull() {
		return false
	}
	q.array[q.rear] = val
	q.rear = (q.rear + 1) % q.size
	return true
}

func (q *queue) dequeue() (int, bool) {
	if q.IsEmpty() {
		return -1, false
	}

	val := q.array[q.front]
	q.front = (q.front + 1) % q.size
	return val, true
}

func (q *queue) getfront() int {
	return q.array[q.front]
}

/**
注意事项：
https://zhuanlan.zhihu.com/p/366159235
数据入队时 end + 1，end 可能数组长度 n，这里有一个小技巧 (end + 1) % n。
数据出队时 start + 1，start 可能超过数组长度 n，也可以使用小技巧 (end + 1) % n。
队尾的数据在 end - 1，但是 end 可能会为 0，这样也会出错，可以使用小技巧 (end - 1 + n) % n。

作者：wendraw
链接：https://leetcode.cn/problems/design-circular-queue/solution/by-wendraw-e6pq/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
**/

func (q *queue) GetRear() int {
	index := (q.rear - q.front + q.size) % q.size
	return q.array[index]
}

func (q *queue) Display() {

}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{size: k, queue: &queue{array: make([]int, k), front: 0, rear: 0}}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.array
}

func (this *MyCircularDeque) DeleteFront() bool {

}

func (this *MyCircularDeque) DeleteLast() bool {

}

func (this *MyCircularDeque) GetFront() int {

}

func (this *MyCircularDeque) GetRear() int {

}

func (this *MyCircularDeque) IsEmpty() bool {

}

func (this *MyCircularDeque) IsFull() bool {

}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */

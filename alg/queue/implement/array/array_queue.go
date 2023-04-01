package array

/**
使用数组模拟循环队列
front:最新插入的数据的index
rear:最后一个元素的下一个index
mod操作主要保证数据索引不越界,比如数据有10个,索引0-9
存储下一个元素rear+1=9+1=10 越界 10 %10 =0 会回到第一个位置
**/

type MyCircularQueue struct {
	capacity int
	data     []int
	front    int
	rear     int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{capacity: k + 1, data: make([]int, k+1), front: 0, rear: 0}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.data[this.rear] = value
	this.rear = (this.rear + 1) % this.capacity
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.front = (this.front + 1) % this.capacity
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.front]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	idx := (this.rear - 1 + this.capacity) % this.capacity
	return this.data[idx]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.front == this.rear
}

func (this *MyCircularQueue) IsFull() bool {
	return this.front == (this.rear+1)%this.capacity
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */

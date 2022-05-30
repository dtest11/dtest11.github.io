package queue

type Queue struct {
	data   []interface{}
	length int
}

func NewQueue() *Queue {
	return &Queue{
		data:   nil,
		length: 0,
	}
}
func (q *Queue) Add(ele interface{}) {
	q.data = append(q.data, ele)
	q.length++
}

func (q *Queue) Dequeue() {
	q.length--
	q.data = q.data[0:]
}

func (q *Queue) Peek() interface{} {
	return q.data[0]
}

func (q *Queue) Empty() bool {
	return q.length == 0
}

func (q *Queue) Data() []interface{} {
	return q.data
}

func (q *Queue) Len() int {
	return q.length
}

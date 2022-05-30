package queue

type Queue[T any] struct {
	elements []T
}

func NewQueue[T any](arr ...T) Queue[T] {
	queue := Queue[T]{elements: make([]T, 0, 20)}
	queue.elements = append(queue.elements, arr...)
	return queue
}

func (q *Queue[T]) PushBack(value T) {
	q.elements = append(q.elements, value)
}

// 从队列头部取出并从删除对应数据。
// 如果队列为空，返回类型对应零值和fasle,否则为front值和true。
func (q *Queue[T]) PopFront() (T, bool) {
	var value T
	if len(q.elements) == 0 {
		return value, false
	}
	value = q.elements[0]
	q.elements = q.elements[1:]
	return value, true
}

func (q Queue[T]) Size() int {
	return len(q.elements)
}

func (q Queue[T]) Empty() bool {
	return len(q.elements) == 0
}

// 如果队列为空，返回类型对应零值和fasle,否则为front值和true。
func (q Queue[T]) Front() (res T, exist bool) {
	if len(q.elements) == 0 {
		return
	}
	return q.elements[0], true
}

// 如果队列为空，返回类型对应零值和fasle,否则为end值和true。
func (q Queue[T]) Back() (res T, exist bool) {
	if len(q.elements) == 0 {
		return
	}
	return q.elements[len(q.elements)-1], true
}

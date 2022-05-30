package queue

import "testing"

func TestPushBackInt(t *testing.T) {
	// queue := NewQueue[int](20, 13)
	queue := NewQueue(20, 13)
	queue.PushBack(26)
	if !(queue.elements[0] == 20 && queue.elements[1] == 13 && queue.elements[2] == 26) {
		t.Error("queue push back int data error.")
	}
}

func TestPushBackFloat64(t *testing.T) {
	queue := NewQueue[float64]()
	queue.PushBack(3.14)
	queue.PushBack(1.15)
	queue.PushBack(2.97)
	if !(queue.elements[0] == 3.14 && queue.elements[1] == 1.15 && queue.elements[2] == 2.97) {
		t.Error("queue push back int data error.")
	}
}

func TestPopFrontInt(t *testing.T) {
	queue := NewQueue(1, 2, 3, 4, 5)
	front, exist := queue.PopFront()
	if front != 1 || exist != true {
		t.Error("queue pop front int data error.")
	}
}

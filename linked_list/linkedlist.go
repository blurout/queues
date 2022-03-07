package linked_list

import "fmt"

type MyQueue struct {
	data int
	next *MyQueue
}

func Make_Queue() MyQueue {
	var queue MyQueue = MyQueue{data: -1, next: nil}
	return queue
}

func (q *MyQueue) Enqueue(x int) {
	if q.data == -1 {
		q.data = x
		return
	}
	var newq MyQueue = MyQueue{data: x, next: nil}
	for i := q; i != nil; i = i.next {
		if i.next == nil {
			i.next = &newq
			return
		}
	}
}

func (q *MyQueue) Dequeue() int {
	if q.next == nil {
		p := q.data
		q.data = 0
		return p
	}
	r := q.data
	*q = *q.next
	return r
}

func (q *MyQueue) Peek() int {
	return q.data
}

func (q *MyQueue) Empty() bool {
	return q.data == 0
}

func (q *MyQueue) Print() {
	curr := q
	for curr != nil {
		fmt.Print(curr.data," ")
		curr = curr.next
	}
}
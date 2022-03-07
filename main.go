package main

import (
	"queues/array"
	"queues/linked_list"
)

func main() {
	q := array.Make_Queue()
	for i := 0; i < 5; i++ {
		q.Enqueue(i)
	}
	q.Print()
	q.Dequeue()
	q.Print()

	q2 := linked_list.Make_Queue()
	for i := 0; i < 5; i++ {
		q2.Enqueue(i)
	}
	q2.Print()
}

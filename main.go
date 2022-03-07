package main

import (
	"queues/array"
)

func main() {
	q := array.Make_Queue()
	for i := 0; i < 5; i++ {
		q.Enqueue(i)
	}
	q.Print()
	q.Dequeue()
	q.Print()
}

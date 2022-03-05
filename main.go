package main

import(
	"fmt"
)

// Queue represents a queue using slice
type Queue struct {
	items []int
}

// Enqueue adds a value at the end

func (q *Queue) Enqueue(n int) {
	q.items = append(q.items, n)
}

// Dequeue removes first value FIFO

func (q *Queue) Dequeue() int {
	remove := q.items[0]
	q.items = q.items[:1]
	return remove
}

func main(){
	q := Queue{}
	fmt.Println(q)
	for i := 0; i < 5; i++ {
		q.Enqueue(i)
	}
	fmt.Println(q)
	q.Dequeue()	
	q.Dequeue()
	q.Dequeue()
	fmt.Println(q)
}

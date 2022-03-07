package array

import(
	"fmt"
)

// Queue represents a queue using slice
type Queue struct {
	items []int
}
// queue contructor called to make new queue object
func Make_Queue() Queue {
	q := Queue{}
	return q
}

// Enqueue adds a value at the end

func (q *Queue) Enqueue(n int) {
	q.items = append(q.items, n)
}

// Dequeue removes first value FIFO

func (q *Queue) Dequeue() int {
	remove := q.items[0]
	q.items = q.items[1:]
	return remove
}

// prints current queue

func (q *Queue) Print() {
	for i := 0; i < len(q.items); i++ {
		fmt.Print(q.items[i]," ")
	}
	fmt.Println()
}
package queue

import "errors"

/*
queue is first in first out ordering.
code one with these operators:
	1. add - add to the back of the queue
	2. remove - remove the first in queue
	3. peek - look at the first object, but don't remove
	4. isempty - is the queue empty
*/

// Queue : Creating a structure which has an array of int32
type Queue struct {
	orderQueue []int32
}

// Add : Inserting element into queue from the end
func (q *Queue) Add(element int32) {
	q.orderQueue = append(q.orderQueue, element)
}

// Remove : Delete element from queue FIFO
func (q *Queue) Remove() {
	// Always check for empty
	if !q.IsEmpty() {
		q.orderQueue = q.orderQueue[1:]
	}
}

// Peek : Access first object
func (q *Queue) Peek() (int32, error) {
	if !q.IsEmpty() {
		return q.orderQueue[0], nil
	}
	return int32(0), errors.New("Empty Queue")
}

// IsEmpty : check for empty queue
func (q *Queue) IsEmpty() bool {
	isEmpty := len(q.orderQueue) < 1
	return isEmpty
}

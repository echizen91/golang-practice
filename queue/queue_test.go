package queue

import (
	"errors"
	"reflect"
	"testing"
)

func TestQueue_Add(t *testing.T) {
	q := Queue{}
	q2 := Queue{}

	compareQ := Queue{[]int32{1, 5, 7}}
	compareQ2 := Queue{[]int32{10374, 197563, 382871}}

	q.Add(1)
	q.Add(5)
	q.Add(7)

	q2.Add(10374)
	q2.Add(197563)
	q2.Add(382871)

	if reflect.DeepEqual(q, compareQ) != true {
		t.Errorf("Adding to Queue Incorrect: %v; want: %v \n", q, compareQ)
	}

	if reflect.DeepEqual(q2, compareQ2) != true {
		t.Errorf("Adding to Queue Incorrect: %v; want: %v \n", q2, compareQ2)
	}
}

func TestQueue_Remove(t *testing.T) {
	q := Queue{[]int32{1, 5, 7, 10}}
	q2 := Queue{[]int32{10374, 197563, 382871}}

	compareQ := Queue{[]int32{5, 7, 10}}
	compareQ2 := Queue{[]int32{197563, 382871}}

	q.Remove()
	q2.Remove()

	if reflect.DeepEqual(q, compareQ) != true {
		t.Errorf("Removing from Queue Incorrect: %v; want: %v \n", q, compareQ)
	}

	if reflect.DeepEqual(q2, compareQ2) != true {
		t.Errorf("Removing from Queue Incorrect: %v; want: %v \n", q2, compareQ2)
	}
}

func TestQueue_Peek(t *testing.T) {
	q := Queue{[]int32{1, 5, 7, 10}}
	q2 := Queue{[]int32{10374, 197563, 382871}}

	compareQ := int32(1)
	compareQ2 := int32(10374)

	qPeek, _ := q.Peek()
	q2Peek, _ := q2.Peek()

	if reflect.DeepEqual(qPeek, compareQ) != true {
		t.Errorf("Peek from Queue Incorrect: %v; want: %v \n", qPeek, compareQ)
	}

	if reflect.DeepEqual(q2Peek, compareQ2) != true {
		t.Errorf("Peek from Queue Incorrect: %v; want: %v \n", q2Peek, compareQ2)
	}

	qEmpty := Queue{}
	_, compareQEmpty := qEmpty.Peek()

	empty := errors.New("Empty Queue")
	// Check for Empty Queue
	if reflect.DeepEqual(compareQEmpty.Error(), empty.Error()) != true {
		t.Errorf("Peek from Queue Incorrect: %v; want: %v \n", compareQEmpty, empty)
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	bFalse := false
	bTrue := true

	q := Queue{[]int32{1}}
	isEmpty := q.IsEmpty()

	// Test 1 have items in it, so False
	if reflect.DeepEqual(isEmpty, bFalse) != true {
		t.Errorf("Is Empty Queue Incorrect: %v; want: %v \n", isEmpty, bFalse)
	}

	q.Remove()
	isEmpty = q.IsEmpty()

	// Test 2 have no items in it, so True
	if reflect.DeepEqual(isEmpty, bTrue) != true {
		t.Errorf("Is Empty Queue Incorrect: %v; want: %v \n", isEmpty, bTrue)
	}
}

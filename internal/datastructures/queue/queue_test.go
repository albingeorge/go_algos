package queue

import "testing"

func TestQueue(t *testing.T) {
	q := Queue{}

	q.Insert(3)
	q.Insert(1)
	q.Insert(6)

	q.Print()
	if q.Size != 3 {
		t.Errorf("size mismatch; expected: %v; received: %v", 3, q.Size)
	}

	val := q.Pop()
	if val != 3 {
		t.Errorf("popped value mismatch; expected: %v; received: %v", 3, val)
	}

	q.Print()

	if q.Size != 2 {
		t.Errorf("size mismatch; expected: %v; received: %v", 2, q.Size)
	}

	q.Pop()

	val = q.Pop()
	if val != 6 {
		t.Errorf("popped value mismatch; expected: %v; received: %v", 6, val)
	}

	if q.Size != 0 {
		t.Errorf("size mismatch; expected: %v; received: %v", 0, q.Size)
	}
}

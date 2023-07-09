package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := Stack{}
	s.Insert(3)
	s.Insert(4)
	s.Insert(1)
	s.Print()
	if s.Size != 3 {
		t.Errorf("expected size: 3; received: %v", s.Size)
	}
	popped := s.Pop()
	if popped != 1 {
		t.Errorf("expected popped: 1; received: %v", popped)
	}
	s.Pop()
	popped = s.Pop()
	if popped != 3 {
		t.Errorf("expected popped: 3; received: %v", popped)
	}
	s.Print()
	if s.Size != 0 {
		t.Errorf("expected size: 0; received: %v", s.Size)
	}
}

package queue

import "fmt"

type QuenueNode struct {
	Next *QuenueNode
	Data int
}

type Queue struct {
	Root *QuenueNode
	Size int
}

func (q *Queue) Insert(data int) {
	node := QuenueNode{
		Data: data,
	}

	if q.Root == nil {
		q.Root = &node
	} else {
		curr := q.Root
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = &node
	}

	q.Size++
}

func (q *Queue) Pop() int {
	var val int
	if q.Root == nil {
		panic("out of range")
	} else if q.Size == 1 {
		val = q.Root.Data
		q.Root = nil
	} else {
		val = q.Root.Data
		q.Root = q.Root.Next
	}
	q.Size--
	return val
}

func (q *Queue) Print() {
	curr := q.Root

	for curr != nil {
		fmt.Printf("%v ", curr.Data)
		curr = curr.Next
	}
	fmt.Println()
}

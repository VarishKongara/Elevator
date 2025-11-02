package main

import (
	"container/heap"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}

type FloorRequest struct {
	Floor int
}

type ElevatorStatus int

const (
	Idle ElevatorStatus = iota
	MovingUp
	MovingDown
)

type Elevator struct {
	ID        int
	Status    ElevatorStatus
	UpQueue   *MinHeap
	DownQueue *MaxHeap
	Location  int
}

func NewElevator(id int) *Elevator {
	up := &MinHeap{}
	down := &MaxHeap{}
	heap.Init(up)
	heap.Init(down)

	return &Elevator{
		ID:        id,
		Status:    Idle,
		UpQueue:   up,
		DownQueue: down,
		Location:  1,
	}
}

func (e *Elevator) Move() {
	switch e.Status {
	case Idle:
		if e.UpQueue.Len() > 0 {
			e.Status = MovingUp
		} else if e.DownQueue.Len() > 0 {
			e.Status = MovingDown
		}
	case MovingUp:
		if e.UpQueue.Len() > 0 {
			nextFloor := (*e.UpQueue)[0]
			if nextFloor == e.Location {
				heap.Pop(e.UpQueue)
			} else {
				e.Location++
			}
		} else if e.DownQueue.Len() > 0 {
			e.Status = MovingDown
		} else {
			e.Status = Idle
		}
	case MovingDown:
		if e.DownQueue.Len() > 0 {
			nextFloor := (*e.DownQueue)[0]
			if nextFloor == e.Location {
				heap.Pop(e.DownQueue)
			} else {
				e.Location--
			}
		} else if e.UpQueue.Len() > 0 {
			e.Status = MovingUp
		} else {
			e.Status = Idle
		}
	}
}

func (e *Elevator) AddRequest(floor int) {
	if e.Status == Idle {
		if floor > e.Location {
			heap.Push(e.UpQueue, floor)
		} else if floor < e.Location {
			heap.Push(e.DownQueue, floor)
		}
	} else {
		switch e.Status {
		case MovingUp:
			if floor >= e.Location {
				heap.Push(e.UpQueue, floor)
			} else {
				heap.Push(e.DownQueue, floor)
			}
		case MovingDown:
			if floor <= e.Location {
				heap.Push(e.DownQueue, floor)
			} else {
				heap.Push(e.UpQueue, floor)
			}
		}
	}
}
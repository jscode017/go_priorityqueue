package pqueue

import (
	"container/heap"
)

type Item struct {
	Value    interface{}
	Priority int64
	Index    int
}

type PriorityQueue []*Item

func New(cap int64) PriorityQueue {
	return make(PriorityQueue, 0, cap)
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) ReSize(newCap int) {
	pqLen := len(*pq)
	newPq := make(PriorityQueue, pqLen, newCap)
	copy(newPq, *pq)
	*pq = newPq
}

func (pq *PriorityQueue) Push(x interface{}) {
	index := len(*pq)
	pqCap := cap(*pq)
	if index >= pqCap {
		pq.ReSize(pqCap * 2)
	}

	item := x.(*Item)
	(*item).Index = index
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	pqLen := len(*pq)
	pqCap := cap(*pq)

	item := (*pq)[pqLen-1]
	(*item).Index = -1
	(*pq) = (*pq)[:pqLen-1]

	if pqLen-1 < int(pqCap/2) {
		pq.ReSize(int(pqCap / 2))
	}

	return (*item)
}

func (pq *PriorityQueue) Update(index int, priority int64, value interface{}) {
	item := (*pq)[index]
	item.Value = value
	item.Priority = priority

	heap.Fix(pq, index)
}

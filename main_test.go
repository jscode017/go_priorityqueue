package pqueue

import (
	"container/heap"
	"testing"
)

func TestPQ_PushAndPop(t *testing.T) {
	pqCap := 50
	pq := New(50)
	for i := 0; i <= 10; i++ {
		heap.Push(&pq, &Item{Value: i, Priority: int64(i)})
	}
	for i := 70; i >= 51; i-- {
		heap.Push(&pq, &Item{Value: i, Priority: int64(i)})
	}
	for i := 11; i <= 50; i++ {
		heap.Push(&pq, &Item{Value: i, Priority: int64(i)})
	}
	if len(pq) != 71 {
		t.Logf("length not equal want %d got %d\n", 71, len(pq))
		t.Fail()
	}
	if cap(pq) != pqCap*2 {

		t.Logf("capacity not equal want %d got %d\n", pqCap*2, cap(pq))
		t.Fail()
	}
	for i := 70; i >= 0; i-- {
		ele := heap.Pop(&pq)
		if ele.(Item).Priority != int64(i) {
			t.Logf("element not equal want %d get %d\n", i, int(ele.(Item).Priority))
			t.Fail()
		}
	}
	if len(pq) != 0 {
		t.Logf("length not equal want %d got %d\n", 0, len(pq))
		t.Fail()
	}

}

func TestPQ_Update(t *testing.T) {

	pq := New(3)
	for i := 0; i <= 2; i++ {
		heap.Push(&pq, &Item{Value: i, Priority: int64(i)})
	}
	pq.Update(0, 5, 5)
	ele := heap.Pop(&pq)
	if ele.(Item).Priority != int64(5) {
		t.Logf("update fail\n")
		t.Fail()
	}
}

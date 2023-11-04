package y2558

import (
	"container/heap"
	"math"
)

func pickGifts(gifts []int, k int) int64 {
	q := &pq{}
	for _, gift := range gifts {
		q.Push(gift)
	}
	heap.Init(q)
	for k > 0 {
		x := heap.Pop(q).(int)
		heap.Push(q, int(math.Floor(math.Sqrt(float64(x)))))
		k--
	}

	var res int64
	for q.Len() > 0 {
		res += int64(q.Pop().(int))
	}
	return res
}

type pq []int

func (q pq) Len() int {
	return len(q)
}

func (q pq) Less(i, j int) bool {
	return q[i] > q[j]
}

func (q pq) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *pq) Push(v interface{}) {
	*q = append(*q, v.(int))
}

func (q *pq) Pop() interface{} {
	n := len(*q)
	res := (*q)[n-1]
	*q = (*q)[0 : n-1]
	return res
}

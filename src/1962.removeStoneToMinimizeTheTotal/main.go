package main

import (
	"container/heap"
	"fmt"
)

func main() {
	piles := []int{5, 4, 9}
	k := 2
	sum := minStoneSum(piles, k)
	fmt.Println(sum)
}

func minStoneSum(piles []int, k int) int {
	// math.Min[int](k, len(piles))
	cnt := k
	if len(piles) < cnt {
		cnt = len(piles)
	}
	x := h(piles)
	heap.Init(&x)
	for i := 0; i < k; i++ {
		pop := heap.Pop(&x)
		remove := pop.(int) / 2
		heap.Push(&x, pop.(int)-remove)
	}
	result := 0
	for _, v := range x {
		result += v
	}
	return result
}

type h []int

func (receiver h) Less(i, j int) bool {
	return receiver[i] > receiver[j]
}

func (receiver h) Swap(i, j int) {
	receiver[i], receiver[j] = receiver[j], receiver[i]
}

func (receiver *h) Push(x any) {
	*receiver = append(*receiver, x.(int))
}

func (receiver *h) Pop() any {
	old := *receiver
	n := len(old)
	x := old[n-1]
	*receiver = old[0 : n-1]
	return x
}

func (receiver h) Len() int {
	return len(receiver)
}

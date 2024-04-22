package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"os"
)

type Pair struct {
	First  int
	Second int
}

type Heap []Pair

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].Second > h[j].Second }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}
func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	p := old[n-1]
	*h = old[:n-1]
	return p
}

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i])
		}

		mp := make(map[int]int)
		for _, v := range arr {
			mp[v]++
		}

		pq := Heap{}
		for k, v := range mp {
			heap.Push(&pq, Pair{k, v})
		}

		size := n
		for len(pq) >= 2 {
			t1 := heap.Pop(&pq).(Pair)
			t2 := heap.Pop(&pq).(Pair)
			c1, c2 := t1.Second, t2.Second
			c1--
			c2--
			size -= 2
			if c1 > 0 {
				heap.Push(&pq, Pair{
					First: t1.First,
					Second: c1,
				})
			}
			if c2 > 0 {
				heap.Push(&pq, Pair{
					First: t2.First,
					Second: c2,
				})
			}
		}
		fmt.Fprintln(out, size)
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}

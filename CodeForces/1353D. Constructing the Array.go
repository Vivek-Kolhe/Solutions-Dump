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

func (h Heap) Len() int      { return len(h) }
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h Heap) Less(i, j int) bool {
	if h[i].Second-h[i].First+1 == h[j].Second-h[j].First+1 {
		return h[i].First < h[j].First
	}
	return h[i].Second-h[i].First+1 > h[j].Second-h[j].First+1
}
func (h *Heap) Push(x any) { *h = append(*h, x.(Pair)) }
func (h *Heap) Pop() any {
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

		pq := Heap{}
		ans := make([]int, n)
		curr := 1
		heap.Push(&pq, Pair{0, n-1})
		for len(pq) > 0 {
			top := heap.Pop(&pq).(Pair)
			l, r := top.First, top.Second
			if l > r {
				continue
			}
			mid := (l + r) >> 1
			ans[mid] = curr
			heap.Push(&pq, Pair{l, mid-1})
			heap.Push(&pq, Pair{mid+1, r})
			curr++
		}
		
		for _, v := range ans {
			fmt.Fprint(out, v, " ")
		}
		fmt.Fprintln(out)
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}

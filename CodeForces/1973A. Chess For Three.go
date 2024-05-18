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

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h Heap) Less(i, j int) bool { return h[i] > h[j] }
func (h *Heap) Push(x any)        { *h = append(*h, x.(int)) }
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
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)

		if a == b && b == c && c == a && a == 0 {
			fmt.Fprintln(out, 0)
			continue
		}
		if (a+b+c)%2 == 1 {
			fmt.Fprintln(out, -1)
			continue
		}

		pq := Heap{}
		if a > 0 {
			heap.Push(&pq, a)
		}
		if b > 0 {
			heap.Push(&pq, b)
		}
		if c > 0 {
			heap.Push(&pq, c)
		}

		var ans int
		for len(pq) > 1 {
			t1, t2 := heap.Pop(&pq).(int), heap.Pop(&pq).(int)
			ans++; t1--; t2--
			if t1 > 0 {
				heap.Push(&pq, t1)
			}
			if t2 > 0 {
				heap.Push(&pq, t2)
			}
		}
		fmt.Fprintln(out, ans)
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}

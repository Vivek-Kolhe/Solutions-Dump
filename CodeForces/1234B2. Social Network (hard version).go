package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
)

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k int
	fmt.Fscan(in, &n, &k)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	q := list.New()
	seen := make(map[int]bool)
	for i := 0; i < n; i++ {
		if _, ok := seen[arr[i]]; !ok {
			q.PushFront(arr[i])
			seen[arr[i]] = true
		}
		for q.Len() > k {
			back := q.Back().Value.(int)
			delete(seen, back)
			q.Remove(q.Back())
		}
	}
	fmt.Fprintln(out, q.Len())
	for e := q.Front(); e != nil; e = e.Next() {
		fmt.Fprint(out, e.Value, " ")
	}
	fmt.Fprintln(out)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}

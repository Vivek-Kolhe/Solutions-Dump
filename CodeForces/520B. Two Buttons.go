package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
)

type Pair struct {
	First  int
	Second int
}

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	visited := make([]bool, int(1e5)+1)
	q := list.New()
	q.PushBack(Pair{0, n})
	visited[n] = true

	for q.Len() > 0 {
		top := q.Front().Value.(Pair)
		q.Remove(q.Front())
		moves, num := top.First, top.Second

		if num == m {
			fmt.Fprintln(out, moves)
			return
		}

		newNum := 2 * num
		if newNum >= 0 && newNum < int(1e5) && !visited[newNum] {
			q.PushBack(Pair{moves+1, newNum})
			visited[newNum] = true
		}
		newNum = num - 1
		if newNum >= 0 && newNum < int(1e5) && !visited[newNum] {
			q.PushBack(Pair{moves+1, newNum})
			visited[newNum] = true
		}
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}

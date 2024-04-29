package main

import (
	"bufio"
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

	var mini, maxi int
	
	if m >= (n+1)/2 {
		mini = n
	} else {
		nodes, edges := n, m
		for nodes > 0 && edges > 0 {
			if nodes > 1 {
				mini += 2
				nodes -= 2
				edges--
			} else {
				nodes--
				edges--
			}
		}
	}

	nodes, edges := n, m
	itr := 1
	if nodes >= 2 && edges > 0 {
		maxi += 2
		nodes -= 2
		edges -= itr
		itr++
	}
	for nodes > 0 && edges > 0 {
		nodes--
		maxi++
		edges -= itr
		itr++
	}

	fmt.Fprintln(out, n-mini, n-maxi)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}

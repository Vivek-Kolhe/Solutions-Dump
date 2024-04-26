package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, u, v int
	fmt.Fscan(in, &n, &m)
	adj := make([][]int, n+1)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &u, &v)
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	inDeg := make([]int, n+1)
	for i := 1; i <= n; i++ {
		for _, v := range adj[i] {
			inDeg[v]++
		}
	}
	
	var one, two, more int
	for i := 1; i <= n; i++ {
		if inDeg[i] == 1 {
			one++
		} else if inDeg[i] == 2 {
			two++
		} else if inDeg[i] > 2 {
			more += inDeg[i]
		}
	}

	if one == 2 && two == n-2 {
		fmt.Fprintln(out, "bus topology")
		return
	}
	if one == 0 && two == n {
		fmt.Fprintln(out, "ring topology")
		return
	}
	if one == n-1 && more == one {
		fmt.Fprintln(out, "star topology")
		return
	}
	fmt.Fprintln(out, "unknown topology")
}

func main() {
	Solve(os.Stdin, os.Stdout)
}

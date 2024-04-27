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

	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
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

		var one, root int
		for i := 1; i <= n; i++ {
			if inDeg[i] == 1 {
				one++
			}
		}
		root = n - one - 1
		one /= root
		fmt.Fprintln(out, root, one)
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}

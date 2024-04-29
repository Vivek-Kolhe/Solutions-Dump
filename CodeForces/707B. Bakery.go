package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	adj := make([][]Pair, n+1)
	for i := 0; i < m; i++ {
		var u, v, wt int
		fmt.Fscan(in, &u, &v, &wt)
		adj[u] = append(adj[u], Pair{v, wt})
		adj[v] = append(adj[v], Pair{u, wt})
	}
	mp := make(map[int]bool)
	flour := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(in, &flour[i])
		mp[flour[i]] = true
	}

	ans := math.MaxInt
	for _, v := range flour {
		for _, wt := range adj[v] {
			if _, ok := mp[wt.First]; ok {
				continue
			}
			ans = min(ans, wt.Second)
		}
	}

	if ans == math.MaxInt {
		fmt.Fprintln(out, -1)
		return
	}
	fmt.Fprintln(out, ans)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Pair struct {
	First  string
	Second string
}

func Solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	var u, v, x string
	fmt.Fscan(in, &n)
	posts := make([]Pair, n)
	mp := make(map[string]int)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &u, &x, &v)
		u = strings.ToLower(u)
		v = strings.ToLower(v)
		posts[i] = Pair{u, v}
		mp[u] = 1
		mp[v] = 1
	}

	for _, p := range posts {
		mp[p.First] = mp[p.Second] + 1
	}

	var ans int
	for _, v := range mp {
		ans = max(ans, v)
	}
	fmt.Fprintln(out, ans)
}

func main() {
	Solve(os.Stdin, os.Stdout)
}

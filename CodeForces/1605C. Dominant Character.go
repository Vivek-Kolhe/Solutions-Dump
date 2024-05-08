package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

type Pair struct {
	First  int
	Second int
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
		var s string
		fmt.Fscan(in, &s)

		possible := []string{"aa", "aba", "aca", "abca", "acba", "abbacca", "accabba"}
		mp := make(map[string]bool)
		for _, p := range possible {
			mp[p] = false
		}

		for _, p := range possible {
			if strings.Contains(s, p) {
				mp[p] = true
			}
		}

		ans := math.MaxInt
		for k, v := range mp {
			if v {
				ans = min(ans, len(k))
			}
		}
		if ans == math.MaxInt {
			ans = -1
		}
		fmt.Fprintln(out, ans)
	}
}

func main() {
	Solve(os.Stdin, os.Stdout)
}
